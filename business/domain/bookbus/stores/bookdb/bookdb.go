package bookdb

import (
	"context"

	"github.com/bentenison/microservice/business/domain/bookbus"
	"github.com/bentenison/microservice/business/sdk/sqldb"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	log *logger.CustomLogger
	db  *sqlx.DB
}

func NewStore(log *logger.CustomLogger, db *sqlx.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

func (s *Store) NewWithTx(tx sqldb.TransactionManager) (bookbus.Storer, error) {

	store := Store{
		log: s.log,
		db:  s.db,
	}
	return &store, nil
}

func (s *Store) Create(ctx context.Context, bks bookbus.Book) error {
	query := `INSERT INTO books (title, author_id, category_id, isbn, published_date, available_copies, tags, image_url,created_at,updated_at) 
          VALUES (:title, :author_id, :category_id, :isbn, :published_date, :available_copies, :tags, :image_url,:created_at,:updated_at)`
	_, err := s.db.NamedExec(query, &bks)
	if err != nil {
		s.log.Errorc(ctx, "error while creating book", map[string]interface{}{
			"error": err,
		})
		return err
	}
	return nil
}
func (s *Store) Update(ctx context.Context, bks bookbus.UpdateBook) error {
	query := `UPDATE books SET available_copies = :available_copies, tags = :tags WHERE book_id = :book_id`
	_, err := s.db.NamedExec(query, &bks)
	if err != nil {
		s.log.Errorc(ctx, "error while creating book", map[string]interface{}{
			"error": err,
		})
		return err
	}
	return nil
}
func (s *Store) Delete(ctx context.Context, bks bookbus.Book) error {
	query := `DELETE FROM books WHERE book_id = :book_id`
	_, err := s.db.NamedExec(query, &bks)
	if err != nil {
		s.log.Errorc(ctx, "error while creating book", map[string]interface{}{
			"error": err,
		})
		return err
	}
	return nil
}

func (s *Store) Query(ctx context.Context) ([]bookbus.Book, error) {
	query := `SELECT title, author_id, category_id, isbn, published_date, available_copies, tags, image_url,created_at,updated_at FROM books`
	var books []bookbus.Book
	err := sqldb.NamedQuerySlice(ctx, s.log, s.db, query, nil, &books)
	if err != nil {
		s.log.Errorc(ctx, "error while querying all books", map[string]interface{}{
			"error": err,
		})
		return books, err
	}
	return books, nil
}

//	func (s *Store) Count(ctx context.Context) (int, error) {
//		return 0, nil
//	}
func (s *Store) QueryByID(ctx context.Context, BookId int) (bookbus.Book, error) {
	return bookbus.Book{}, nil
}
func (s *Store) QueryByUserID(ctx context.Context, BookId int) ([]bookbus.Book, error) {
	return []bookbus.Book{}, nil
}
