package bookdb

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/bentenison/microservice/business/domain/bookbus"
)

func applyFilter(filter bookbus.QueryFilter, data map[string]any, buf *bytes.Buffer) {
	var wc []string
	if filter.BookId != nil {
		data["book_id"] = *filter.BookId
		wc = append(wc, "book_id = :book_id")
	}
	if filter.Title != nil {
		data["title"] = fmt.Sprintf("%%%s%%", *filter.Title)
		wc = append(wc, "title LIKE :title")
	}
	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}
