package bookbus

import "github.com/bentenison/microservice/business/sdk/order"

// DefaultOrderBy represents the default way we sort.
var DefaultOrderBy = order.NewBy(OrderByBookID, order.ASC)

// Set of fields that the results can be ordered by.
const (
	OrderByBookID        = "book_id"
	OrderByCreatedAt     = "created_at"
	OrderByUpdatedAt     = "updated_at"
	OrderByPublishedDate = "published_date"
	OrderByTitle         = "title"
)
