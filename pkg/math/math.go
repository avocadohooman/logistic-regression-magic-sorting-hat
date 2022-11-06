package math

import (
	"fmt"
	"math"

	utils "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/utils"
)

func Count(columns utils.Column) int {
	return len(columns)
}

func Mean(column utils.Column) float64 {
	var retVal float64

	for _, value := range column {
		retVal += value
	}

	retVal = (retVal / float64(Count(column)))

	return retVal
}

func Std(columns utils.Column) float64 {
	var retVal float64

	mean := Mean(columns)
	for _, value := range columns {
		retVal += (value - mean) * (value - mean)
	}
	retVal = retVal / (float64)(len(columns)-1)
	retVal = math.Sqrt(retVal)

	return retVal
}

func Min() {}

func Max() {}

func Medium() {}

func Percentile(percentile int) float64 {
	return 0.00
}
