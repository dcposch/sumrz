package main
import (
    "fmt"
    "math"
    "strconv"
    "strings"
)

func (s *NumberStats) init() {
    s.Max = math.Inf(-1)
    s.Min = math.Inf(1)
}

func (s *NumberStats) update(value string) {
    cleanVal := strings.Replace(value, ",", "", -1)
    f, err := strconv.ParseFloat(cleanVal, 64)
    if err != nil {
        return
    }
    if math.IsInf(f,0) || math.IsNaN(f) {
        return
    }
    _, err = strconv.Atoi(cleanVal)
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
        ret = "real numbers, "
    } else {
        ret = "integers, "
    }

    if s.NumValidNums == 0 {
        ret = "not numeric"
    } else if s.Min == s.Max {
        if s.IsFloat {
            ret += fmt.Sprintf("always %e", s.Min)
        } else {
            ret += fmt.Sprintf("integers, always %d", int64(s.Min))
        }
    } else {
        avg := s.Sum / float64(s.NumValidNums)
        stdev := math.Sqrt((s.SumSquares / float64(s.NumValidNums)) - avg*avg)

        if s.IsFloat {
            ret += fmt.Sprintf("min %e max %e avg %e stdev %e",
                s.Min, s.Max, avg, stdev)
        } else {
            ret += fmt.Sprintf("min %d max %d avg %f stdev %f",
                int64(s.Min), int64(s.Max), avg, stdev)
        }
    }
    return ret
}

