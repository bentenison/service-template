package bookapp

import (
	"github.com/bentenison/microservice/business/domain/bookbus"
	"github.com/bentenison/microservice/business/sdk/order"
)

var defaultOrderBy = order.NewBy("book_id", order.ASC)
var orderByFields = map[string]string{
	"book_id":        bookbus.OrderByBookID,
	"title":          bookbus.OrderByTitle,
	"created_at":     bookbus.OrderByCreatedAt,
	"published_date": bookbus.OrderByPublishedDate,
	"updated_at":     bookbus.OrderByUpdatedAt,
}
