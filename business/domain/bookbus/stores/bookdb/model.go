package bookdb

import (
	"time"
)

type Book struct {
	BookID          int       `db:"book_id" json:"book_id,omitempty"`
	Title           string    `db:"title" json:"title,omitempty"`
	AuthorID        int       `db:"author_id" json:"author_id,omitempty"`
	CategoryID      int       `db:"category_id" json:"category_id,omitempty"`
	ISBN            string    `db:"isbn" json:"isbn,omitempty"`
	PublishedDate   time.Time `db:"published_date" json:"published_date,omitempty"`
	AvailableCopies int       `db:"available_copies" json:"available_copies,omitempty"`
	Tags            []string  `db:"tags" json:"tags,omitempty"`
	CreatedAt       time.Time `db:"created_at" json:"created_at,omitempty"`
	ImageUrl        string    `db:"image_url" json:"image_url,omitempty"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

// func toBusBook() bookbus.Book {

// }
