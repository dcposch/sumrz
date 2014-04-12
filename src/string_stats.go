package main
import (
    "math/rand"
)

const maxCounts = 1000
const numTopCounts = 10
const numSample = 10

func (s *StringStats) init() {
    s.Counts = make(map[string]int64)
}

// Updates the top value estimate and the random sample.
// `numNotBlank` is 1 for the first update and increments each time
func (s *StringStats) update(value string, numNotBlank int64) {
    // Constant-space streaming estimator for the top item counts
    // http://www.cs.berkeley.edu/~satishr/cs270/sp11/rough-notes/Streaming-two.pdf
    s.Counts[value]++
    if len(s.Counts) > maxCounts {
        s.IsEstimate = true
        var valsToDelete []string
        for val,count := range(s.Counts) {
            if count == 1 {
                valsToDelete = append(valsToDelete, val)
            } else {
                s.Counts[val]--
            }
        }
        for _, val := range(valsToDelete) {
            delete(s.Counts, val)
        }
    }

    // Random sample
    if (numNotBlank <= numSample) {
        s.Sample = append(s.Sample, value)
    } else if rand.Int63n(numNotBlank) < numSample {
        evictIndex := rand.Intn(numSample)
        s.Sample[evictIndex] = value
    }
}

