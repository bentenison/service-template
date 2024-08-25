package bookbus

import "time"

type NewBook struct {
	Title           string    `json:"title"`
	ISBN            string    `json:"isbn"`
	PublishedDate   time.Time `json:"published_date"`
	AvailableCopies int       `json:"available_copies"`
	Tags            []string  `json:"tags"`
}
type Book struct {
	BookID          int       `json:"book_id,omitempty"`
	Title           string    `json:"title,omitempty"`
	AuthorID        int       `json:"author_id,omitempty"`
	CategoryID      int       `json:"category_id,omitempty"`
	ISBN            string    `json:"isbn,omitempty"`
	PublishedDate   time.Time `json:"published_date,omitempty"`
	AvailableCopies int       `json:"available_copies,omitempty"`
	Tags            []string  `json:"tags,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	ImageUrl        string    `json:"image_url,omitempty"`
}
type UpdateBook struct {
	BookID          int      `json:"book_id,omitempty"`
	Tags            []string `json:"tags"`
	AvailableCopies int      `json:"available_copies"`
}
