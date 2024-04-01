package grpc_service

import (
	"context"

	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
)

func (i *Implementation) Ping(ctx context.Context, _ *inquiry.PingRequest) (*inquiry.PingResponse, error) {
	text, err := i.inquiryService.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &inquiry.PingResponse{Status: text}, nil
}
