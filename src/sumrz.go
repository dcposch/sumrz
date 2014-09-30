package main
import (
    "encoding/csv"
    "flag"
    "fmt"
    "io"
    "os"
    //"runtime/pprof"
)


func main() {
    //f,_ := os.Create("sumrz.prof")
    //pprof.StartCPUProfile(f)
    //defer pprof.StopCPUProfile()

    var delim = '\t'
    flag.Parse()
    args := flag.Args()
    if len(args) > 1 {
        printUsageExit()
    }
    if len(args) == 1 {
        if len(args[0]) != 1 {
            printUsageExit()
        }
        delim = rune(args[0][0])
    }

    var stats TableStats
    err := readCsvAndComputeStats(os.Stdin, &stats, delim)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    } else {
        fmt.Println(stats.String())
    }
}

func printUsageExit() {
    fmt.Fprint(os.Stderr, "Usage: sumrz < file.csv\n" +
                          "   or: sumrz '\t' < file.txt")
    os.Exit(1)
}

func readCsvAndComputeStats(reader io.Reader, stats *TableStats, delim rune) error {
    csvReader := csv.NewReader(reader)
    csvReader.LazyQuotes = true
    csvReader.Comma = delim
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
    if err != io.EOF {
        return fmt.Errorf("Failed on line %d: %v\n", lineNum, err)
    }
    return nil
}

