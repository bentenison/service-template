package mid

import (
	"context"
	"net/http"

	"github.com/bentenison/microservice/foundation/web"
	"github.com/google/uuid"
)

func TraceTdMiddleware(next http.Handler) web.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.NewString()
		}

		ctx := context.WithValue(r.Context(), traceKey, correlationID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
}
