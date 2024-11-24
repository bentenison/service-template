package authdb

import (
	"context"
	"database/sql"

	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/domain/authbus"
	"github.com/bentenison/microservice/foundation/logger"
)

type Store struct {
	log *logger.CustomLogger
	ds  mux.DataSource
}

func NewStore(log *logger.CustomLogger, ds mux.DataSource) *Store {
	return &Store{
		log: log,
		ds:  ds,
	}
}

func (s *Store) CreateUser(ctx context.Context, user *authbus.User) (string, error) {
	//create user query
	query := `INSERT INTO users (id, username, email, password_hash, first_name, last_name, role, created_at, updated_at)
		VALUES (:id, :username, :email, :password_hash, :first_name, :last_name, :role, :created_at, :updated_at)
		RETURNING id`
	var lastInsertedID string
	stmt, err := s.ds.SQL.PrepareNamed(query)
	if err != nil {
		s.log.Errorc(ctx, "error while adding user", map[string]interface{}{
			"error": err.Error(),
		})
		return "", err
	}
	err = stmt.Get(&lastInsertedID, user)
	if err != nil {
		s.log.Errorc(ctx, "error while adding user", map[string]interface{}{
			"error": err.Error(),
		})
		return "", err
	}
	return lastInsertedID, nil

}
func (s *Store) GetUser(ctx context.Context, u string) (*authbus.User, error) {
	//create user query
	query := `
		SELECT id, username, email, password_hash, first_name, last_name, role, created_at, updated_at
		FROM users
		WHERE username = $1
		LIMIT 1`

	var user UserDB

	err := s.ds.SQL.Get(&user, query, u)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	busUser := toBusUser(&user)
	return busUser, nil
}
func (s *Store) ListUsers(ctx context.Context) ([]*authbus.User, error) {
	//create user query
	query := `
		SELECT id, username, email, password_hash, first_name, last_name, role, created_at, updated_at
		FROM users
	`
	var users []UserDB

	err := s.ds.SQL.Select(&users, query)
	if err != nil {
		return nil, err
	}
	busUsers := toBustUsers(users)
	return busUsers, nil
}
