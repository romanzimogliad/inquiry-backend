package grpc_server

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"

	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Login creates new token
func (i *Implementation) Login(ctx context.Context, request *inquiry.LoginRequest) (*inquiry.LoginResponse, error) {
	response, err := i.inquiryService.Login(ctx, &domain.LoginRequest{
		Login:    request.GetLogin(),
		Password: request.GetPassword(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error in Login: %v", err)
	}
	return &inquiry.LoginResponse{
		Token: response.Token,
	}, nil
}
