package main
import (
    "fmt"
    "math"
    "strconv"
)

func (s *NumberStats) update(value string) {
    f, err := strconv.ParseFloat(value, 64)
    if err != nil && err != strconv.ErrRange {
        return
    }
    _, err = strconv.Atoi(value)
    if err != nil {
        // Value is not an integer
        s.IsFloat = true
    }
    s.NumValidNums++
    s.Min = math.Min(s.Min, f)
    s.Max = math.Max(s.Max, f)
    s.Sum += f
    s.SumSquares += f*f
}

func (s *NumberStats) String() string {
    var ret string
    if s.IsFloat {
        ret = ""
    } else {
        ret = "integers, "
    }

    if s.NumValidNums == 0 {
        ret = "not numeric"
    } else if s.Min == s.Max {
        if s.IsFloat {
            ret += fmt.Sprintf("always %f", s.Min)
        } else {
            ret += fmt.Sprintf("integers, always %d", int64(s.Min))
        }
    } else {
        avg := s.Sum / float64(s.NumValidNums)
        stdev := math.Sqrt((s.SumSquares - avg*avg) / float64(s.NumValidNums))

        if s.IsFloat {
            ret += fmt.Sprintf("min %f max %f avg %f stdev %f",
                s.Min, s.Max, avg, stdev)
        } else {
            ret += fmt.Sprintf("min %d max %d avg %f stdev %f",
                int64(s.Min), int64(s.Max), avg, stdev)
        }
    }
    return ret
}

