package mid

import (
	"context"
	"net/http"

	"github.com/bentenison/microservice/business/sdk/sqldb"
)

type TransactionMiddleware struct {
	tm sqldb.TransactionManager
}

func NewTransactionMiddleware(tm sqldb.TransactionManager) *TransactionMiddleware {
	return &TransactionMiddleware{tm: tm}
}

func (tmw *TransactionMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Start a new transaction
		tx, err := tmw.tm.BeginTx(ctx)
		if err != nil {
			http.Error(w, "Could not start transaction", http.StatusInternalServerError)
			return
		}

		// Set the transaction in the repository
		// tmw.repo.SetTx(tx)

		// Create a new context with the transaction and pass it to the handler
		ctx = context.WithValue(ctx, "tx", tx)
		r = r.WithContext(ctx)

		// Call the next handler in the chain
		defer func() {
			if err := recover(); err != nil {
				_ = tmw.tm.RollbackTx(tx) // Rollback on panic
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
		}()
		next.ServeHTTP(w, r)

		// If no error, commit the transaction
		if err := tmw.tm.CommitTx(tx); err != nil {
			http.Error(w, "Could not commit transaction", http.StatusInternalServerError)
			return
		}
	})
}
