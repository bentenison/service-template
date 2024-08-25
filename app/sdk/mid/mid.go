package mid

import "context"

type ctxKey int

const (
	claimKey ctxKey = iota + 1
	userIDKey
	userKey
	bookKey
	trKey
	traceKey
)

func GetTraceId(ctx context.Context) string {
	return ctx.Value(traceKey).(string)
}
