package page

import (
	"fmt"
	"strconv"
)

type Page struct {
	number int
	rows   int
}

func Parse(rowsPerPage string, pageNumber string) (Page, error) {
	number := 1
	if pageNumber != "" {
		var err error
		number, err = strconv.Atoi(pageNumber)
		if err != nil {
			return Page{}, fmt.Errorf("page coversion: %w", err)
		}
	}
	rows := 10
	if rowsPerPage != "" {
		var err error
		rows, err = strconv.Atoi(rowsPerPage)
		if err != nil {
			return Page{}, fmt.Errorf("rows coversion: %w", err)
		}
	}
	//page and rows validation
	if number <= 0 {
		return Page{}, fmt.Errorf("page number cannot be zero or negative")
	}
	if rows <= 0 {
		return Page{}, fmt.Errorf("rows per page cannot be zero or negative")
	}
	if rows > 100 {
		return Page{}, fmt.Errorf("rows per page cannot be graetor than 100")
	}
	p := Page{
		number: number,
		rows:   rows,
	}
	return p, nil
}

func (p Page) PageNumber() int {
	return p.number
}
func (p Page) RowsPerPage() int {
	return p.rows
}
