package math

import (
	"errors"
	"strconv"

	"fmt"
)

type Columns []float64

func ToFloat64(columns []string) ([]float64, error) {
	newColumns := make(Columns, len(columns))

	for i, value := range columns {
		if parsed, err := strconv.ParseFloat(value, 64); err != nil {
			return nil, errors.New(fmt.Sprintf("Cannot parse %v to float64", value))
		} else {
			newColumns[i] = parsed
		}
	}
	return newColumns, nil
}

func Count(columns Columns) int {
	return len(columns)
}

func Mean(columns Columns) float64 {
	var total float64

	for _, value := range columns {
		total=+value
	}

	return (total / float64(Count(columns)))
}

func Std() {}

func Min() {}

func Max() {}

func TopPercentage() {}

func CardinalityFloat64(column Columns) int {
	set := make(map[float64]int)

	for _, value := range column {
		set[value] = 1
	}

	return len(set)
}

func CardinalityString(column []string) int {
	set := make(map[string]int)

	for _, value := range column {
		set[value] = 1
	}

	return len(set)
}
