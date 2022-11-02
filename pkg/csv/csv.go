package csv

import (
	"encoding/csv"
	"os"

	errors "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/errors"
)

type Csv struct {
	headers []string
	// columnar
	data [][]string
	columns [][]string
}

func (c *Csv) GetColumns(index int) ([]string, error){
	if index > len(c.headers) {
		return nil, errors.New("Invalid header index")
	}

	return c.columns[index], nil
}

func readCsv(file *os.File) *csv.Reader {
	return csv.NewReader(file)
}

func ParseCsv(csv_file *os.File) *Csv {
	csv_reader := readCsv(csv_file)
	headers, err := csv_reader.Read()
	errors.DieIfErr(err)

	data := make([][]string, 0)
	columns := make([][]string, len(headers))

	for row, err := csv_reader.Read(); err == nil; row, err = csv_reader.Read(){
		data = append(data, row)

		columnIndex := 0
		for index, field := range row {
			if index < 2 {
				continue
			}

			if len(columns) == 0 {
				columns[columnIndex] = make([]string, 0)
			}
			columns[columnIndex] = append(columns[columnIndex], field)
			columnIndex++
		}
	}

	return &Csv{
		headers: headers,
		data: data,
		columns: columns,
	}
}

