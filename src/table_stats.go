package main
import (
    "fmt"
    "os"
    "strings"
)

func (s *TableStats) init(headers []string) error {
    s.Headers = make([]string, len(headers))
    s.Fields = make(map[string]*FieldStats)
    for i,header := range(headers) {
        header = strings.TrimSpace(header)
        if _,ok := s.Fields[header]; ok {
            fmt.Fprintf(os.Stderr, "Duplicate header '%s'\n", header)
        }
        stats := new(FieldStats)
        stats.init()
        s.Headers[i] = header
        s.Fields[header] = stats
    }
    return nil
}

func (s *TableStats) update(values []string) error {
    if len(values) != len(s.Headers) {
        return fmt.Errorf("Found %d values but %d headers",
            len(values), len(s.Headers))
    }
    for i,header := range(s.Headers) {
        field := s.Fields[header]
        if field != nil {
            field.update(values[i])
        }
    }
    s.NumRows++
    return nil
}

func (s *TableStats) String() string {
    tableStr := fmt.Sprintf("Table stats - %d columns, %d rows\n"+
                            "=================================",
                            len(s.Headers), s.NumRows)
    var fieldStrs []string
    for _,header := range(s.Headers) {
        var headerStr string
        if header == "" {
            headerStr = "(blank header)"
        } else {
            headerStr = "'"+header+"'"
        }
        fieldStr := s.Fields[header].String()
        fieldStrs = append(fieldStrs, fmt.Sprintf("%s %s",
            headerStr, fieldStr))
    }
    return tableStr + "\n\n" + strings.Join(fieldStrs, "\n\n")
}

