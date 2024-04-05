package grpc_server

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
	"github.com/romanzimoglyad/inquiry-backend/internal/grpc_server/mappings"
	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetLesson gets lesson
func (i *Implementation) GetLesson(ctx context.Context, request *inquiry.GetLessonRequest) (*inquiry.GetLessonResponse, error) {
	lesson, err := i.inquiryService.GetLesson(ctx, &domain.GetLessonsRequest{
		UserId: request.GetUserId(),
		Id:     request.GetId(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error in ListLessons: %v", err)
	}

	return &inquiry.GetLessonResponse{
		Lesson: mappings.FormLesson(lesson),
	}, nil
}
