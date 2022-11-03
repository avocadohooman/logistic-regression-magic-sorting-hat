package main

import (
	"fmt"
	csv "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/csv"
	errors "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/errors"
	math "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/math"
	utils "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/utils"
	"log"
	"os"
)

type Stats struct {
	FeatureName string
	Count int
	Mean float64
	Std float64
}


func main() {
	args := os.Args
	dateColumn := 4
	// handColumn := 3
	skipColumns := []int{0, 1, 2, 3, 5}

	if len(args) < 2 {
		log.Fatal("File path required")
	}
	filep := args[1]
	csv, err := csv.ParseCsv(filep)
	errors.DieIfErr(err)

	var stats []Stats
	for i := 0; i < csv.GetHeaderCount(); i++ {
		if utils.Contains(skipColumns, i) {
			continue
		}

		column, err := csv.GetColumns(i)
		errors.DieIfErr(err)

		if i == dateColumn {
			column, err = utils.ParseDate(column)
			errors.DieIfErr(err)
		}

		parseColumn, err := math.ToFloat64(column)

		stat := Stats{
			FeatureName: csv.GetHeader(i),
			Count: math.Count(parseColumn),
			Mean: math.Mean(parseColumn),
			Std: math.Std(parseColumn),
		}
		stats = append(stats, stat)
	}

	fmt.Println(stats)
}
