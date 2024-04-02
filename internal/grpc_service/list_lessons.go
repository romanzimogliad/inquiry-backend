package grpc_service

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"

	"github.com/romanzimoglyad/inquiry-backend/internal/grpc_service/mappings"
	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListLessons list lessons
func (i *Implementation) ListLessons(ctx context.Context, request *inquiry.ListLessonsRequest) (*inquiry.ListLessonsResponse, error) {
	lessons, err := i.inquiryService.ListLessons(ctx, &domain.ListLessonsRequest{UserId: request.GetUserId(), Filter: domain.Filter{
		SubjectId: request.GetFilter().GetSubjectId(),
		ConceptId: request.GetFilter().GetConceptId(),
		UnitId:    request.GetFilter().GetUnitId(),
		SkillId:   request.GetFilter().GetSkillId(),
	}})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error in ListLessons: %v", err)
	}

	resp := make([]*inquiry.Lesson, len(lessons))

	for k, v := range lessons {
		resp[k] = mappings.FormLesson(v)
	}
	return &inquiry.ListLessonsResponse{
		Lessons: resp,
	}, nil
}
