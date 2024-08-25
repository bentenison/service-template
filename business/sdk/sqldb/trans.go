package sqldb

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type TransactionManager interface {
	BeginTx(ctx context.Context) (*sql.Tx, error)
	CommitTx(tx *sql.Tx) error
	RollbackTx(tx *sql.Tx) error
}

type PostgresManager struct {
	db *sqlx.DB
}

func NewTransactionManager(db *sqlx.DB) TransactionManager {
	return &PostgresManager{
		db: db,
	}
}
func (tm *PostgresManager) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return tm.db.BeginTx(ctx, &sql.TxOptions{})
}

func (tm *PostgresManager) CommitTx(tx *sql.Tx) error {
	return tx.Commit()
}

func (tm *PostgresManager) RollbackTx(tx *sql.Tx) error {
	return tx.Rollback()
}

// func GetExtContext(tx CommitRollbacker) (sqlx.ExtContext, error) {
// 	ec, ok := tx.(sqlx.ExtContext)
// 	if !ok {
// 		return nil, fmt.Errorf("Transactor(%T) not of a type *sql.Tx", tx)
// 	}

// 	return ec, nil
// }
