package authdb

import (
	"database/sql"

	"github.com/bentenison/microservice/business/domain/authbus"
)

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDB struct {
	ID           sql.NullString `json:"id,omitempty" db:"id"`
	Username     sql.NullString `json:"username,omitempty" db:"username"`
	Email        sql.NullString `json:"email,omitempty" db:"email"`
	PasswordHash sql.NullString `json:"password_hash,omitempty" db:"password_hash"`
	FirstName    sql.NullString `json:"first_name,omitempty" db:"first_name"`
	LastName     sql.NullString `json:"last_name,omitempty" db:"last_name"`
	Role         sql.NullString `json:"role,omitempty" db:"role"`
	CreatedAt    sql.NullTime   `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at,omitempty" db:"updated_at"`
}

func toBusUser(user *UserDB) *authbus.User {
	var u authbus.User
	u.Email = user.Email.String
	u.Username = user.Username.String
	u.PasswordHash = user.PasswordHash.String
	u.FirstName = user.FirstName.String
	u.LastName = user.LastName.String
	u.Role = user.Role.String
	u.ID = user.ID.String
	return &u
}

func toBustUsers(users []UserDB) []*authbus.User {
	busUsers := make([]*authbus.User, len(users))
	for _, v := range users {
		u := toBusUser(&v)
		busUsers = append(busUsers, u)
	}
	return busUsers
}
