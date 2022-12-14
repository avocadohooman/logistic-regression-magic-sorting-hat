package math

import (
	"errors"
	"math"
	"sort"

	"github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/models/stats"
)

func Count(columns stats.Column) int {
	return len(columns)
}

func Mean(column stats.Column) float64 {
	var retVal float64

	for _, value := range column {
		retVal += value
	}

	retVal = (retVal / float64(Count(column)))

	return retVal
}

func Std(columns stats.Column) float64 {
	var retVal float64

	mean := Mean(columns)
	for _, value := range columns {
		retVal += (value - mean) * (value - mean)
	}

	retVal = retVal / (float64)(len(columns)-1)
	retVal = math.Sqrt(retVal)

	return retVal
}

func Min(values []float64) float64 {
	if len(values) == 0 {
		return math.NaN()
	}

	minValue := values[0]

	if len(values) == 1 {
		return minValue
	}

	for _, value := range values {
		if float64(value) < minValue {
			minValue = float64(value)
		}
	}
	return float64(minValue)
}

func Max(values []float64) float64 {
	if len(values) == 0 {
		return math.NaN()
	}

	maxValue := values[0]

	if len(values) == 1 {
		return maxValue
	}

	for _, value := range values {
		if float64(value) > maxValue {
			maxValue = float64(value)
		}
	}
	return float64(maxValue)
}

func Percentile(input []float64, percentile float64) (float64, error) {
	if len(input) == 0 {
		return math.NaN(), errors.New("Invalid input")
	}
	if len(input) == 1 {
		return input[0], nil
	}
	if percentile <= 0 || percentile > 100 {
		return math.NaN(), errors.New("Percentile range needs to be between 0 - 100")
	}

	sortedCopy := sortSlice(input)

	// Multiply percent by length of input
	// n = (P/100) x N
	index := (percentile / 100) * float64(len(sortedCopy))

	// Check if the index is a whole number
	if index == float64(int64(index)) {
		// Convert float to int
		i := int(index)
		// Find&return the value at the index
		return sortedCopy[i-1], nil
	} else if index > 1 {
		// Convert float to int via truncation
		i := int(index)
		// Find the average of the index and following values
		meanPercentile := Mean([]float64{sortedCopy[i-1], sortedCopy[i]})
		return meanPercentile, nil
	}
	return math.NaN(), errors.New("Invalid input")
}

func sortSlice(input []float64) []float64 {
	copiedSlice := copySlice(input)
	sort.Float64s(copiedSlice)
	return copiedSlice
}

func copySlice(input []float64) []float64 {
	newSlice := make([]float64, len(input))
	copy(newSlice, input)
	return newSlice
}
