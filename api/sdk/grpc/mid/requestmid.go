package mid

import (
	"context"

	"github.com/bentenison/microservice/foundation/logger"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// type contextKey string

// const requestIDKey contextKey = "tracectx"

type RequestIDMiddleware struct {
	logger *logger.CustomLogger
}

// NewRequestIDMiddleware constructor to create a new instance of the middleware with a logger
func NewRequestIDMiddleware(logger *logger.CustomLogger) *RequestIDMiddleware {
	return &RequestIDMiddleware{logger: logger}
}
func (r *RequestIDMiddleware) UnaryRequestIDInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Retrieve request ID from incoming metadata or create a new one
		md, ok := metadata.FromIncomingContext(ctx)
		var requestID string

		if ok {
			ids := md.Get("request-id")
			if len(ids) > 0 {
				requestID = ids[0]
			}
		}

		if requestID == "" {
			// Generate a new UUID if no request ID was provided
			requestID = uuid.New().String()
		}

		// Log the request ID
		// log.Printf("Request ID: %s", requestID)

		// Add the request ID to the context
		ctx = context.WithValue(ctx, "tracectx", requestID)

		// Proceed with the request
		return handler(ctx, req)
	}
}
func (r *RequestIDMiddleware) StreamRequestIDInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		// Retrieve request ID from incoming metadata or create a new one
		md, ok := metadata.FromIncomingContext(ss.Context())
		var requestID string

		if ok {
			ids := md.Get("request-id")
			if len(ids) > 0 {
				requestID = ids[0]
			}
		}

		if requestID == "" {
			// Generate a new UUID if no request ID was provided
			requestID = uuid.New().String()
		}

		// Log the request ID
		// log.Printf("Stream Request ID: %s", requestID)

		// Add the request ID to the context
		wrappedStream := &serverStreamWithContext{
			ServerStream: ss,
			ctx:          context.WithValue(ss.Context(), "tracectx", requestID),
		}

		// Proceed with the stream
		return handler(srv, wrappedStream)
	}
}

// Custom server stream wrapper to allow context modifications
type serverStreamWithContext struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *serverStreamWithContext) Context() context.Context {
	return w.ctx
}
