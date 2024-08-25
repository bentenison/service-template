package order

import (
	"fmt"
	"strings"
)

const (
	ASC  = "ASC"
	DESC = "DESC"
)

var directions = map[string]string{
	ASC:  "ASC",
	DESC: "DESC",
}

type By struct {
	Field     string
	Direction string
}

func NewBy(field string, direction string) By {
	if _, exists := directions[direction]; !exists {
		return By{
			Field:     field,
			Direction: ASC,
		}
	}
	return By{
		Field:     field,
		Direction: direction,
	}
}

// Parse constructs a By value by parsing a string in the form of
// "field,direction" ie "user_id,ASC".
func Parse(fieldsMappings map[string]string, orderBy string, defaultOrder By) (By, error) {
	if orderBy == "" {
		return defaultOrder, nil
	}
	//"field,direction" ie "user_id,ASC". slice the string to diffrent values
	orderParts := strings.Split(orderBy, ",")
	orgFieldName := strings.TrimSpace(orderParts[0])
	fieldName, exists := fieldsMappings[orgFieldName]
	if !exists {
		return By{}, fmt.Errorf("unknown order: %s", orgFieldName)
	}
	switch len(orderParts) {
	case 1:
		// only field is there so use default direction ASC
		return NewBy(fieldName, ASC), nil
	case 2:

		//parse direction and apply it
		direction := strings.TrimSpace(orderParts[1])
		if _, exists := directions[direction]; !exists {
			return By{}, fmt.Errorf("unknown direction: %s", direction)
		}

		return NewBy(fieldName, direction), nil
	default:
		return By{}, fmt.Errorf("unknown order: %s", orderBy)
	}
}
