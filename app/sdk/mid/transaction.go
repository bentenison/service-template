package mid

import (
	"net/http"

	"github.com/bentenison/microservice/business/sdk/sqldb"
	"github.com/gin-gonic/gin"
)

type TransactionMiddleware struct {
	tm sqldb.TransactionManager
}

func NewTransactionMiddleware(tm sqldb.TransactionManager) *TransactionMiddleware {
	return &TransactionMiddleware{tm: tm}
}

func (tmw *TransactionMiddleware) TransactionManager() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ctx := r.Context()

		// Start a new transaction
		tx, err := tmw.tm.BeginTx(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Could not start transaction")
			return
		}

		// Set the transaction in the repository
		// tmw.repo.SetTx(tx)

		// Create a new context with the transaction and pass it to the handler
		c.Set(trKey, tx)
		// r = r.WithContext(ctx)

		// Call the next handler in the chain
		defer func() {
			if err := recover(); err != nil {
				_ = tmw.tm.RollbackTx(tx) // Rollback on panic
				c.JSON(http.StatusInternalServerError, "Internal server error")
				return
			}
		}()
		c.Next()

		// If no error, commit the transaction
		if err := tmw.tm.CommitTx(tx); err != nil {
			c.JSON(http.StatusInternalServerError, "Could not commit transaction")
			// http.Errorc(w,  http.StatusInternalServerError)
			return
		}
	}
}
