package utils

import (
	"net/url"
	"strconv"
)

type Filter struct {
	Field    string
	Value    string
	Operator string
}

func ParseFilters(query url.Values) ([]Filter, error) {
	filters := []Filter{}

	for key, values := range query {
		if len(values) > 0 {
			filter := Filter{
				Field:    key,
				Value:    values[0],
				Operator: "eq",
			}

			if len(values) > 1 {
				operator, err := strconv.Atoi(values[1])
				if err != nil {
					return nil, err
				}

				switch operator {
				case 1:
					filter.Operator = "eq"
				case 2:
					filter.Operator = "neq"
				case 3:
					filter.Operator = "gt"
				case 4:
					filter.Operator = "gte"
				case 5:
					filter.Operator = "lt"
				case 6:
					filter.Operator = "lte"
				default:
					return nil, err
				}
			}

			filters = append(filters, filter)
		}
	}

	return filters, nil
}