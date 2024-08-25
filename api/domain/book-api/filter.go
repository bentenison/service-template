package bookapi

import (
	"net/http"

	"github.com/bentenison/microservice/app/domain/bookapp"
)

func parseQueryParams(r *http.Request) (bookapp.QueryParam, error) {
	values := r.URL.Query()

	filter := bookapp.QueryParam{
		Page:          values.Get("page"),
		Rows:          values.Get("row"),
		OrderBy:       values.Get("orderBy"),
		BookId:        values.Get("book_id"),
		Title:         values.Get("title"),
		PublishedDate: values.Get("published_date"),
		CreatedAt:     values.Get("created_at"),
		UpdatedAt:     values.Get("updated_at"),
	}

	return filter, nil
}
