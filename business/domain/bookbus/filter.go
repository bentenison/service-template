package bookbus

import (
	"time"
)

// QueryFilter holds the available fields a query can be filtered on.
// We are using pointer semantics because the With API mutates the value.
type QueryFilter struct {
	BookId           *int
	ISBN             *string
	Title            *string
	StartCreatedDate *time.Time
	EndCreatedDate   *time.Time
}
