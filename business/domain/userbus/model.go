package userbus

import "time"

type User struct {
	UserID      int       `db:"user_id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Email       string    `db:"email"`
	PhoneNumber string    `db:"phone_number"`
	Password    string    `db:"password"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type Lease struct {
	LeaseID    int       `db:"lease_id"`
	UserID     int       `db:"user_id"`
	BookID     int       `db:"book_id"`
	LeaseDate  time.Time `db:"lease_date"`
	ReturnDate time.Time `db:"return_date"`
	IsReturned bool      `db:"is_returned"`
}
