package bookbus

import (
	"context"

	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/business/sdk/sqldb"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/jmoiron/sqlx"
)

type Storer interface {
	NewWithTx(tx sqldb.TransactionManager) (Storer, error)
	Create(ctx context.Context, bks Book) error
	Update(ctx context.Context, bks UpdateBook) error
	Delete(ctx context.Context, bks Book) error
	Query(ctx context.Context) ([]Book, error)
	// Count(ctx context.Context) (int, error)
	QueryByID(ctx context.Context, BookId int) (Book, error)
	QueryByUserID(ctx context.Context, BookId int) ([]Book, error)
}

type Business struct {
	log      *logger.CustomLogger
	db       *sqlx.DB
	delegate *delegate.Delegate
	storer   Storer
}

func NewBusiness(log *logger.CustomLogger, db *sqlx.DB, delegate *delegate.Delegate, storer Storer) *Business {
	return &Business{
		log:      log,
		db:       db,
		delegate: delegate,
		storer:   storer,
	}
}
func (b *Business) NewWithTx(tx sqldb.TransactionManager) (Storer, error) {
	return b.storer.NewWithTx(tx)
}

func (b *Business) Create(ctx context.Context, bks Book) error {

	return nil
}
func (b *Business) Update(ctx context.Context, bks Book) error {

	return nil
}
func (b *Business) Delete(ctx context.Context, bks Book) error {

	return nil
}

func (b *Business) Query(ctx context.Context) ([]Book, error) {
	return []Book{}, nil
}

//	func (b *Business) Count(ctx context.Context) (int, error) {
//		return 0, nil
//	}
func (b *Business) QueryByID(ctx context.Context, BookId int) (Book, error) {
	return Book{}, nil
}
func (b *Business) QueryByUserID(ctx context.Context, BookId int) ([]Book, error) {
	return []Book{}, nil
}
