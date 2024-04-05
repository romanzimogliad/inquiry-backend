package grpc_server

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/romanzimoglyad/inquiry-backend/internal/grpc_server/mappings"

	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
)

// CreateLesson creates new lesson
func (i *Implementation) CreateLesson(ctx context.Context, request *inquiry.CreateLessonRequest) (*inquiry.CreateLessonResponse, error) {
	logger.Info().Msgf("Req : %s", request)
	id, err := i.inquiryService.CreateLesson(ctx, mappings.ToLesson(request))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error in CreateLesson: %v", err)
	}
	return &inquiry.CreateLessonResponse{
		Id: id,
	}, nil
}
