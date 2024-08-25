package query

import "github.com/bentenison/microservice/business/sdk/page"

type Result[T any] struct {
	Items       []T `json:"items,omitempty"`
	Total       int `json:"total,omitempty"`
	PageNumber  int `json:"page_number,omitempty"`
	RowsPerPage int `json:"rows_per_page,omitempty"`
}

func NewResult[T any](items []T, total int, page page.Page) Result[T] {
	return Result[T]{
		Items:       items,
		Total:       total,
		PageNumber:  page.PageNumber(),
		RowsPerPage: page.RowsPerPage(),
	}
}
