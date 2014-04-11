package main

import (
    "fmt"
    "sort"
    "strings"
)

func (s *FieldStats) init(){
    s.StrStats.init()
}

func (s *FieldStats) update(value string){
    if value == "" {
        return
    }
    s.NumNotBlank++
    s.StrStats.update(value, s.NumNotBlank)
    s.NumStats.update(value)
}

// Show a summary of the field
func (s *FieldStats) String() string {
    // One-line summary for special cases
    if s.NumNotBlank == 0 {
        return "all blank"
    }
    if len(s.StrStats.Counts) == 1 && !s.StrStats.IsEstimate {
        return fmt.Sprintf("always '%s'", getOnlyValue(s.StrStats.Counts))
    }

    // Multiline summary
    var lines []string
    s.maybeAddNumStats(&lines)
    topCounts := s.getTopCounts()
    totalTopCount := s.maybeShowTopCounts(&lines, topCounts)
    s.maybeAddSample(&lines, totalTopCount)
    return strings.Join(lines, "\n")
}

func (s *FieldStats) maybeAddNumStats(lines *[]string) {
    if s.NumStats.NumValidNums == s.NumNotBlank {
        *lines = append(*lines, s.NumStats.String())
    } else if s.NumStats.NumValidNums > 0 {
        str := fmt.Sprintf("%d/%d numeric: %s",
            s.NumStats.NumValidNums,
            s.NumNotBlank,
            s.NumStats.String())
        *lines = append(*lines, str)
    }
}

func (s *FieldStats) maybeShowTopCounts(lines *[]string, topCounts []fieldCount) int64 {
    tc := topCounts[0].count
    allCounts := len(topCounts) == len(s.StrStats.Counts)
    shouldShowTopN := tc > 10
    exactTopN := !s.StrStats.IsEstimate && shouldShowTopN
    goodEnoughEstimate := tc > (s.NumNotBlank / maxCounts) && shouldShowTopN
    showCounts := allCounts || exactTopN || goodEnoughEstimate
    if allCounts {
        *lines = append(*lines, "All values:")
    } else if exactTopN {
        *lines = append(*lines, "Most common values:")
    } else if goodEnoughEstimate {
        *lines = append(*lines, "Estimated most values. APPROXIMATE counts:")
    }
    totalTopCount := int64(0)
    if showCounts {
        for _,c := range(topCounts) {
            totalTopCount += c.count
            *lines = append(*lines, fmt.Sprintf("%15d '%s'", c.count, c.value))
        }
    }
    return totalTopCount
}

func (s *FieldStats) maybeAddSample(lines *[]string, totalTopCount int64) {
    // If we showed no counts or only a minority of the values, 
    // then augment it with a random sample
    if totalTopCount <= s.NumNotBlank / 2 {
        *lines = append(*lines, "Random sample:")
        for _,val := range(s.StrStats.Sample) {
            *lines = append(*lines, fmt.Sprintf("                '%s'", val))
        }
    }
}

func getOnlyValue(counts map[string]int64) string {
    for val,_ := range(counts) {
        return val
    }
    panic("getOnlyValue() called on empty map")
}


type fieldCount struct {
    value string
    count int64
}

type fieldCountList []fieldCount

func (f fieldCountList) Len() int {
    return len(f)
}

func (f fieldCountList) Less(i, j int) bool {
    // Sort descending by count, ascending by value
    if f[i].count == f[j].count {
        return f[i].value < f[j].value
    }
    return f[i].count > f[j].count
}

func (f fieldCountList) Swap(i, j int) {
    f[i], f[j] = f[j], f[i]
}

// Returns the N most common values from s.StrStats, 
// sorted descending order by count, then ascending by value
func (s *FieldStats) getTopCounts() fieldCountList {
    var topCounts fieldCountList
    for val,count := range(s.StrStats.Counts) {
        topCounts = append(topCounts, fieldCount {val,count})
    }
    sort.Sort(topCounts)
    if len(topCounts) > numTopCounts {
        return topCounts[:numTopCounts]
    }
    return topCounts
}
