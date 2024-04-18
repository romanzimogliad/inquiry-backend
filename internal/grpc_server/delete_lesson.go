package grpc_server

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteLesson delete lesson
func (i *Implementation) DeleteLesson(ctx context.Context, request *inquiry.GetLessonRequest) (*inquiry.Empty, error) {
	err := i.inquiryService.DeleteLesson(ctx, &domain.GetLessonsRequest{
		UserId: request.GetUserId(),
		Id:     request.GetId(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error in DeleteLesson: %v", err)
	}

	return &inquiry.Empty{}, nil
}
