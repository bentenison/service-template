package bookdb

import (
	"fmt"

	"github.com/bentenison/microservice/business/domain/bookbus"
	"github.com/bentenison/microservice/business/sdk/order"
)

var orderByFields = map[string]string{
	bookbus.OrderByBookID:        "book_id",
	bookbus.OrderByCreatedAt:     "created_at",
	bookbus.OrderByUpdatedAt:     "updated_at",
	bookbus.OrderByPublishedDate: "published_date",
	bookbus.OrderByTitle:         "title",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
