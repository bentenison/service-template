package authapp

import (
	"time"

	"github.com/bentenison/microservice/business/domain/authbus"
	"github.com/google/uuid"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	PasswordHash string    `json:"password_hash"`
	FirstName    string    `json:"first_name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserPayload struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordHash string `json:"password_hash"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Role         string `json:"role"`
}

func toBusUser(up *UserPayload) *authbus.User {
	var u authbus.User
	u.ID = uuid.NewString()
	u.Username = up.Username
	u.Email = up.Email
	u.Password = up.Password
	u.PasswordHash = up.PasswordHash
	u.FirstName = up.FirstName
	u.LastName = up.LastName
	u.Role = up.Role
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return &u
}
