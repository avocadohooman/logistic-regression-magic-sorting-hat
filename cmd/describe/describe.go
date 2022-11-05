package main

import (
	"fmt"
	"log"
	"os"

	csvService "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/csv"
	errors "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/errors"
	math "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/math"
	models "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/models/stats"
	utils "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/utils"
)

// consts
var dateColumn = 4
var skipColumns = []int{0, 1, 2, 3, 5}

// var handColumn := 3

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("File path required")
	}

	filep := args[1]
	csv, err := csvService.ParseCsv(filep)
	errors.DieIfErr(err)

	stats := models.StatsArray{}
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

		parseColumn, err := utils.ToFloat64(column)
		errors.DieIfErr(err)

		twentyFive, err := math.Percentile(parseColumn, 25)
		errors.DieIfErr(err)

		fifty, err := math.Percentile(parseColumn, 50)
		errors.DieIfErr(err)

		seventyFive, err := math.Percentile(parseColumn, 75)
		errors.DieIfErr(err)

		stat := models.Stats{
			FeatureName: csv.GetHeader(i),
			Count:       math.Count(parseColumn),
			Mean:        math.Mean(parseColumn),
			Std:         math.Std(parseColumn),
			Min:         math.Min(parseColumn),
			Max:         math.Max(parseColumn),
			TwentyFive:  twentyFive,
			Fifty:       fifty,
			SeventyFive: seventyFive,
		}
		stats = append(stats, stat)
	}

	fmt.Println(stats)
	fmt.Println(csvService.CreateCSV(stats))
}
