package authbus

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT secret key (use environment variable or config for production)
var jwtSecret = []byte("super-secret-key")

// Struct to represent claims inside the JWT
type Claims struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID           string    `json:"id,omitempty" db:"id"`
	Username     string    `json:"username,omitempty" db:"username"`
	Email        string    `json:"email,omitempty" db:"email"`
	PasswordHash string    `json:"password_hash,omitempty" db:"password_hash"`
	FirstName    string    `json:"first_name,omitempty" db:"first_name"`
	LastName     string    `json:"last_name,omitempty" db:"last_name"`
	Role         string    `json:"role,omitempty" db:"role"`
	CreatedAt    time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" db:"updated_at"`
	Password     string    `json:"password,omitempty" db:"-"`
}

// func toDBUser(user User) *authdb.UserDB {
// 	var u *authdb.UserDB
// 	u.Email.String = user.Email
// 	u.Username.String = user.Username
// 	u.PasswordHash.String = user.PasswordHash
// 	u.FirstName.String = user.FirstName
// 	u.LastName.String = user.LastName
// 	u.Role.String = user.Role
// 	u.ID.String = user.ID
// 	return u
// }
