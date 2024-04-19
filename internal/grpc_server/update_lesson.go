package grpc_server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/romanzimoglyad/inquiry-backend/internal/grpc_server/mappings"

	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
)

// UpdateLesson updates lesson
func (i *Implementation) UpdateLesson(ctx context.Context, request *inquiry.LessonRequest) (*inquiry.Empty, error) {
	err := i.inquiryService.UpdateLesson(ctx, mappings.ToLesson(request))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error in UpdateLesson: %v", err)
	}
	return &inquiry.Empty{}, nil
}
