package interceptor

import (
	"context"
	"fmt"

	"github.com/romanzimoglyad/inquiry-backend/internal/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthInterceptor Middleware interceptor for authentication and authorization
func AuthInterceptor(mtr map[string]int32) func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		if roleId, ok := mtr[info.FullMethod]; ok {

			// Extract metadata from context
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, status.Errorf(codes.Unauthenticated, "metadata is missing")
			}
			fmt.Println(info.FullMethod)
			// Extract token from metadata
			token := md.Get("authorization")
			if len(token) == 0 {
				return nil, status.Errorf(codes.Unauthenticated, "token is missing")
			}

			// Validate token
			claims, err := auth.ValidateToken(token[0])
			if err != nil {
				return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
			}

			// Check if user has required role(s)

			if roleId < claims.Role {
				return nil, status.Errorf(codes.PermissionDenied, "insufficient permissions")
			}
		}

		// Proceed to the next handler
		return handler(ctx, req)
	}
}
