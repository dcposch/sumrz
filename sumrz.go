package main
import (
    "encoding/csv"
    "fmt"
    "io"
    "os"
    //"runtime/pprof"
)


func readCsvAndComputeStats(reader io.Reader, stats *TableStats){
    //f,_ := os.Create("sumrz.prof")
    //pprof.StartCPUProfile(f)
    //defer pprof.StopCPUProfile()

    csvReader := csv.NewReader(reader)
    csvReader.Comma = '|'
    var err error
    lineNum := 0
    for {
        var values []string
        values, err = csvReader.Read()
        if err != nil {
            break
        }
        lineNum++
        if lineNum == 1 {
            err = stats.init(values)
        } else {
            err = stats.update(values)
        }
        if lineNum % 10000 == 0 || err == io.EOF {
            fmt.Fprintf(os.Stderr, "Read %d lines\n", lineNum)
        }
        if err != nil {
            break
        }
    }
    if(err != io.EOF){
        fmt.Fprintf(os.Stderr, "Failed on line %d: %v\n", lineNum, err)
    }
}

func main() {
    var stats TableStats
    readCsvAndComputeStats(os.Stdin, &stats)
    fmt.Println(stats.String())
}

