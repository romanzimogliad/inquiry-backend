package service

import (
	"context"

	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
)

func (i *Implementation) Ping(_ context.Context, _ *inquiry.PingRequest) (*inquiry.PingResponse, error) {
	return &inquiry.PingResponse{Status: "Ok"}, nil
}
