package main

import (
	"fmt"
	csv "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/csv"
	errors "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/errors"
	math "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/math"
	"log"
	"os"
	"strings"
)

type Stats struct {
	Count float64
}


func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("File path required")
	}
	filep := args[1]
	if strings.HasSuffix(filep, ".csv") == false {
		log.Fatal("File must be an csv")
	}

	csv_file, err := os.Open(filep)
	errors.DieIfErr(err)

	csv := csv.ParseCsv(csv_file)
	column, err := csv.GetColumns(3)
	errors.DieIfErr(err)

	parsedColumn, err := math.ToFloat64(column)
	errors.DieIfErr(err)
	fmt.Printf("count: %v", math.Count(parsedColumn))
}
