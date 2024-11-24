package mid

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"log"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/metadata"
// )

// // Context keys for storing user information
// type AuthInfo struct {
// 	UserID string
// 	Roles  []string
// }

// func validateToken(token string) (*AuthInfo, error) {
// 	// Simulate token validation (JWT validation can be done here)
// 	if token == "valid-token" {
// 		return &AuthInfo{
// 			UserID: "user-123",
// 			Roles:  []string{"user", "admin"}, // Example roles
// 		}, nil
// 	}
// 	return nil, errors.New("invalid token")
// }

// // Unary interceptor for authentication
// func UnaryAuthInterceptor(
// 	ctx context.Context,
// 	req interface{},
// 	info *grpc.UnaryServerInfo,
// 	handler grpc.UnaryHandler,
// ) (interface{}, error) {
// 	log.Printf("Unary Auth Interceptor triggered for method: %s", info.FullMethod)

// 	// Extract metadata from context (for example, an authorization token)
// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if !ok {
// 		return nil, fmt.Errorf("missing metadata")
// 	}

// 	// Check if the metadata contains the authorization token
// 	token, ok := md["authorization"]
// 	if !ok || len(token) == 0 {
// 		return nil, fmt.Errorf("authorization token not provided")
// 	}

// 	// Validate the token
// 	authInfo, err := validateToken(token[0])
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid token: %v", err)
// 	}

// 	// Add AuthInfo to the context so it can be accessed in handlers
// 	ctx = context.WithValue(ctx, "authInfo", authInfo)

// 	// Call the handler
// 	return handler(ctx, req)
// }

// // // Example function to check for a specific role
// // func authorize(ctx context.Context, requiredRole string) error {
// // 	// Retrieve AuthInfo from context
// // 	authInfo, ok := ctx.Value("authInfo").(*AuthInfo)
// // 	if !ok {
// // 		return fmt.Errorf("authorization information missing")
// // 	}

// // 	// Check if the user has the required role
// // 	for _, role := range authInfo.Roles {
// // 		if role == requiredRole {
// // 			return nil
// // 		}
// // 	}
// // 	return fmt.Errorf("access denied: user does not have the '%s' role", requiredRole)
// // }
