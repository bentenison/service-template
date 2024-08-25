package mid

import (
	"context"
	"net/http"

	"github.com/bentenison/microservice/foundation/web"
	"github.com/google/uuid"
)

func TraceTdMiddleware(next web.HandlerFunc) web.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) any {
		correlationID := r.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.NewString()
		}

		ctx := context.WithValue(r.Context(), traceKey, correlationID)
		r = r.WithContext(ctx)
		// next.ServeHTTP(w, r)
		return next(w, r)
	}
}
