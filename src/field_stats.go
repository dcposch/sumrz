package main

import (
    "fmt"
    "sort"
    "strings"
)

func (s *FieldStats) init(){
    s.NumStats.init()
    s.StrStats.init()
}

func (s *FieldStats) update(value string){
    if value == "" {
        s.NumBlank++
    } else {
        s.NumNotBlank++
        s.StrStats.update(value, s.NumNotBlank)
        s.NumStats.update(value)
    }
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
    lines := []string{s.getHeaderLine()}
    topCounts := s.getTopCounts()
    totalTopCount := s.maybeShowTopCounts(&lines, topCounts)
    s.maybeAddSample(&lines, totalTopCount)
    return strings.Join(lines, "\n")
}

// Returns descriptions like "8 blanks, 10 text" or "all integers, min 3 max 8 avg 4.50000 stdev 0.50000" 
func (s *FieldStats) getHeaderLine() string {
    var parts []string
    n := s.NumBlank + s.NumNotBlank
    if s.NumBlank > 0 {
        parts = append(parts, fmt.Sprintf("%d blanks", s.NumBlank))
    }
    numText := s.NumNotBlank - s.NumStats.NumValidNums
    if numText == n {
        parts = append(parts, "all text")
    } else if numText > 0 {
        parts = append(parts, fmt.Sprintf("%d text", numText))
    }
    if s.NumStats.NumValidNums == n {
        str := fmt.Sprintf("all %s",
            s.NumStats.String())
        parts = append(parts, str)
    } else if s.NumStats.NumValidNums > 0 {
        str := fmt.Sprintf("%d %s",
            s.NumStats.NumValidNums,
            s.NumStats.String())
        parts = append(parts, str)
    }
    return strings.Join(parts, ", ")
}

func (s *FieldStats) maybeShowTopCounts(lines *[]string, topCounts []fieldCount) int64 {
    tc := topCounts[0].count
    allCounts := len(topCounts) == len(s.StrStats.Counts) && !s.StrStats.IsEstimate
    shouldShowTopN := tc > 10
    exactTopN := !s.StrStats.IsEstimate && shouldShowTopN
    goodEnoughEstimate := tc > (s.NumNotBlank / maxCounts) && shouldShowTopN
    showCounts := allCounts || exactTopN || goodEnoughEstimate
    if allCounts {
        *lines = append(*lines, "All values:")
    } else if exactTopN {
        *lines = append(*lines, "Most common values:")
    } else if goodEnoughEstimate {
        *lines = append(*lines, "Estimated most common values. APPROXIMATE counts:")
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
