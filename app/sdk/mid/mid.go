package mid

import (
	"context"
)

const (
	claimKey  = "claimctx"
	userIDKey = "userIdctx"
	userKey   = "userctx"
	bookKey   = "bookctx"
	trKey     = "trsnctx"
	traceKey  = "tracectx"
)

func GetTraceId(ctx context.Context) any {
	return ctx.Value(traceKey)
}
