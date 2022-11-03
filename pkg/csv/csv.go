package csv

import (
	"encoding/csv"
	"os"
	"strings"

	errors "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/errors"
)

type Csv struct {
	headers []string
	// columnar
	data    [][]string
	columns [][]string
}

func (c *Csv) GetColumns(index int) ([]string, error) {
	if index > len(c.headers) {
		return nil, errors.New("Invalid header index")
	}

	return c.columns[index], nil
}

func (c *Csv) GetHeaderCount() int {
	return len(c.headers)
}

func (c *Csv) GetHeader(index int) string {
	return c.headers[index]
}

func (c *Csv) GetRowCount() int {
	return len(c.data[0])
}

func readCsv(file *os.File) *csv.Reader {
	return csv.NewReader(file)
}

func openCsv(path string) (*os.File, error) {
	if strings.HasSuffix(path, ".csv") == false {
		return nil, errors.New("File must be an csv")
	}

	csv_file, err := os.Open(path)
	errors.DieIfErr(err)

	return csv_file, nil
}

func ParseCsv(path string) (*Csv, error) {
	csv_file, err := openCsv(path)
	if err != nil {
		return nil, err
	}
	csv_reader := readCsv(csv_file)
	headers, err := csv_reader.Read()
	if err != nil {
		return nil, err
	}

	data := make([][]string, 0)
	columns := make([][]string, len(headers))

	for row, err := csv_reader.Read(); err == nil; row, err = csv_reader.Read() {
		data = append(data, row)

		columnIndex := 0
		for _, field := range row {
			if len(columns) == 0 {
				columns[columnIndex] = make([]string, 0)
			}
			columns[columnIndex] = append(columns[columnIndex], field)
			columnIndex++
		}
	}

	return &Csv{
		headers: headers,
		data:    data,
		columns: columns,
	}, nil
}
