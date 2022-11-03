package utils

import (
	"errors"
	"fmt"
	"strconv"
)

type Columns []float64

func ToFloat64(columns []string) ([]float64, error) {
	newColumns := make(Columns, len(columns))

	for i, value := range columns {
		if value == "" {
			newColumns[i] = float64(0.00)
		} else {
			parsed, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("Cannot parse %v to float64", value))
			}
			newColumns[i] = parsed
		}
	}
	return newColumns, nil
}

func ParseDate(column []string) ([]string, error) {
	parsedColumn := make([]string, len(column))

	for i, date := range column {
		if len(date) < 10 {
			return nil, errors.New("Invalid date string")
		}
		parsedColumn[i] = date[5:7]
	}

	return parsedColumn, nil
}

func Contains(elements []int, needle int) bool {
	for _, value := range elements {
		if needle == value {
			return true
		}
	}
	return false
}

// Need generics :(
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
