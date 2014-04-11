package main
import (
    "fmt"
    "strings"
)

func (s *TableStats) init(headers []string) error {
    s.Headers = headers
    s.Fields = make(map[string]*FieldStats)
    for _,header := range(headers) {
        header = strings.TrimSpace(header)
        if _,ok := s.Fields[header]; ok {
            return fmt.Errorf("Duplicate header '%s'", header)
        }
        stats := new(FieldStats)
        stats.StrStats.init()
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
        s.Fields[header].update(values[i])
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
        if strings.Contains(fieldStr, "\n") {
            // Show a paragraph summary of this field
            fieldStrs = append(fieldStrs, fmt.Sprintf("%s\n%s",
                headerStr, fieldStr))
        } else {
            // Show a single-line summary of this field
            fieldStrs = append(fieldStrs, fmt.Sprintf("%s %s",
                headerStr, fieldStr))
        }
    }
    return tableStr + "\n\n" + strings.Join(fieldStrs, "\n\n")
}

