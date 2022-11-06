package utils

import (
	"errors"

	"github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/models/stats"
)

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
func CardinalityFloat64(column stats.Column) int {
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
