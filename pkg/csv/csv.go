package csv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	errors "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/errors"
	"github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/models/stats"
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
	defer csv_file.Close()

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

func CreateCSV(describedStats stats.StatsArray) string {
	b := new(bytes.Buffer)
	writer := csv.NewWriter(b)

	var headers []string
	// headers  = append(headers, "")
	for _, describedStat := range describedStats {
		fmt.Println(describedStat.FeatureName)
		headers = append(headers, describedStat.FeatureName)
	}
	writer.Write(headers)

	// for i = 0; i < 8; i++ {
	// 	for 
	// 	var row []string
	// 	row = populateRow(row, describedStat)
	// 	err := writer.Write(row)
	// 	errors.DieIfErr(err)
	// }
	writer.Flush()
	return strings.Join(strings.Split(b.String(), "\n"), "\n")
}

func populateRow(row []string, describedStat stats.Stats) []string {
	return append(row,
		fmt.Sprintf("%d", describedStat.Count),
		fmt.Sprintf("%.4f", describedStat.Mean),
		fmt.Sprintf("%.4f", describedStat.Std),
		fmt.Sprintf("%.4f", describedStat.Max),
		fmt.Sprintf("%.4f", describedStat.Min),
		fmt.Sprintf("%.4f", describedStat.TwentyFive),
		fmt.Sprintf("%.4f", describedStat.Fifty),
		fmt.Sprintf("%.4f", describedStat.SeventyFive),
	)
}
