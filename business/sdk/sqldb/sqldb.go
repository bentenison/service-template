package sqldb

import (
	"context"
	"database/sql"
	"errors"
	"net/url"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

var (
	ErrDBNotFound        = sql.ErrNoRows
	ErrDBDuplicatedEntry = errors.New("duplicated entry")
	ErrUndefinedTable    = errors.New("undefined table")
)

// Config is the required properties to use the database.
type Config struct {
	User         string
	Password     string
	Host         string
	Name         string
	Schema       string
	MaxIdleConns int
	MaxOpenConns int
	DisableTLS   bool
}

// Open knows how to open a database connection based on the configuration.
// func Open(cfg Config) (*pgx.Conn, error) {
// 	sslMode := "require"
// 	if cfg.DisableTLS {
// 		sslMode = "disable"
// 	}

// 	q := make(url.Values)
// 	q.Set("sslmode", sslMode)
// 	q.Set("timezone", "utc")
// 	if cfg.Schema != "" {
// 		q.Set("search_path", cfg.Schema)
// 	}

// 	u := url.URL{
// 		Scheme:   "postgres",
// 		User:     url.UserPassword(cfg.User, cfg.Password),
// 		Host:     cfg.Host,
// 		Path:     cfg.Name,
// 		RawQuery: q.Encode(),
// 	}

//		conf, err := pgxpool.ParseConfig(u.String())
//		if err != nil {
//			return nil, err
//		}
//		conf.MaxConns = int32(runtime.NumCPU())
//		conf.MinConns = int32(runtime.NumCPU() / 2)
//		db, err := pgxpool.ConnectConfig(context.Background(), conf)
//		if err != nil {
//			return nil, err
//		}
//		// db.
//		// db.(cfg.MaxIdleConns)
//		// db.SetMaxOpenConns(cfg.MaxOpenConns)
//	   db.
//		return db, nil
//	}
func Open(cfg Config) (*sqlx.DB, error) {

	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.User, cfg.Password),
		Host:   cfg.Host,
		Path:   cfg.Name,
		// RawQuery: q.Encode(),
	}

	db, err := sqlx.Open("postgres", u.String())
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	return db, nil
}

// StatusCheck returns nil if it can successfully talk to the database. It
// returns a non-nil error otherwise.
func StatusCheck(ctx context.Context, db *sqlx.DB) error {

	// If the user doesn't give us a deadline set 1 second.
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Second)
		defer cancel()
	}

	for attempts := 1; ; attempts++ {
		if err := db.Ping(); err == nil {
			break
		}

		time.Sleep(time.Duration(attempts) * 100 * time.Millisecond)

		if ctx.Err() != nil {
			return ctx.Err()
		}
	}

	if ctx.Err() != nil {
		return ctx.Err()
	}

	// Run a simple query to determine connectivity.
	// Running this query forces a round trip through the database.
	const q = `SELECT TRUE`
	var tmp bool
	return db.QueryRowContext(ctx, q).Scan(&tmp)
}
