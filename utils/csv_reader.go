package utils

import (
    "encoding/csv"
    "io"
    "os"
)

func ReadCSV(file *os.File) ([][]string, error) {
    reader := csv.NewReader(file)
    var rows [][]string
    for {
        row, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }
        rows = append(rows, row)
    }
    return rows, nil
}
