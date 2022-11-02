package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	errors "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/errors"
	csv "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/csv"
)


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
	fmt.Println(csv.GetColumns(1))
}
