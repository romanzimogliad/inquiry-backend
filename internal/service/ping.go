package service

import (
	"context"
	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
)

func (i *Implementation) Ping(ctx context.Context, req *inquiry.PingRequest) (*inquiry.PingResponse, error) {
	return &inquiry.PingResponse{Status: "Ok"}, nil
}
