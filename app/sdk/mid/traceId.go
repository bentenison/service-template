package mid

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		correlationID := c.GetHeader("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.NewString()
		}
		ctx := context.WithValue(c.Request.Context(), traceKey, correlationID)
		c.Request = c.Request.WithContext(ctx)
		// c.Set(traceKey, correlationID)
		c.Next()
	}
}
