package main

// Stores the stats of each field (column) in a table,
// as well as the number of rows of the whole table
type TableStats struct {
    Headers []string
    Fields map[string]*FieldStats
    NumRows int64
}

// Stores the stats of a single field (aka column)
type FieldStats struct {
    StrStats StringStats
    NumStats NumberStats
    NumNotBlank int64
    NumBlank int64
}

// Stores stats relevant for non-numeric fields.
// For example, a field that represents a category, or a Name field.
// 
// Tracks item frequency with a constant-space streaming algorithm described here: 
// http://www.cs.berkeley.edu/~satishr/cs270/sp11/rough-notes/Streaming-two.pdf
type StringStats struct {
    IsEstimate bool
    Counts map[string]int64
    Sample []string
}

// Stores stats relevant for numeric fields.
type NumberStats struct {
    NumValidNums int64
    Min float64
    Max float64
    Sum float64
    SumSquares float64
    IsFloat bool
}

