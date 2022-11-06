package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	csvService "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/csv"
	errors "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/errors"
	mathService "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/math"
	"github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/models/stats"
	models "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/models/stats"
	"github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/print-data"
	utils "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/utils"
)

// consts
var dateColumn = 4
var skipColumns = []int{0, 1, 2, 3, 5}

// var handColumn := 3

func ToFloat64(column []string) ([]float64, error) {
	newColumns := make(stats.Column, len(column))

	for i, value := range column {
		if value == "" {
			newColumns[i] = math.NaN()
		} else {
			parsed, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("Cannot parse %v to float64", value))
			}
			newColumns[i] = parsed
		}
	}
	for _, value := range newColumns {
		if math.IsNaN(value) {
			value = mathService.Min(newColumns)
		}
		fmt.Println("FLOAT 64", value)
	}
	return newColumns, nil
}

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

		parseColumn, err := ToFloat64(column)
		errors.DieIfErr(err)

		twentyFive, err := mathService.Percentile(parseColumn, 25)
		errors.DieIfErr(err)

		fifty, err := mathService.Percentile(parseColumn, 50)
		errors.DieIfErr(err)

		seventyFive, err := mathService.Percentile(parseColumn, 75)
		errors.DieIfErr(err)

		stat := models.Stats{
			FeatureName: csv.GetHeader(i),
			Count:       mathService.Count(parseColumn),
			Mean:        mathService.Mean(parseColumn),
			Std:         mathService.Std(parseColumn),
			Min:         mathService.Min(parseColumn),
			Max:         mathService.Max(parseColumn),
			TwentyFive:  twentyFive,
			Fifty:       fifty,
			SeventyFive: seventyFive,
		}
		stats = append(stats, stat)
	}
	printdata.PrintData(stats)
}
