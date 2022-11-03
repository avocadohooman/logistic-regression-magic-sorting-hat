package math

import (
	utils "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/utils"
)

func Count(columns utils.Columns) int {
	return len(columns)
}

func Mean(columns utils.Columns) float64 {
	var retVal float64

	for _, value := range columns {
		retVal = +value
	}
	retVal = (retVal / float64(Count(columns)))

	return retVal
}

func Std(columns utils.Columns) float64 {
	var retVal float64

	mean := Mean(columns)
	for _, value := range columns {
		retVal = +(value - mean) * (value - mean)
	}
	retVal = retVal / (float64)(len(columns)-1)

	return retVal
}

func Min() {}

func Max() {}

func Medium() {}

func Percentile(percentile int) float64 {
	return 0.00
}
