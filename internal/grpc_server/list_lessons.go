package grpc_server

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"

	"github.com/romanzimoglyad/inquiry-backend/internal/grpc_server/mappings"
	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListLessons list lessons
func (i *Implementation) ListLessons(ctx context.Context, request *inquiry.ListLessonsRequest) (*inquiry.ListLessonsResponse, error) {
	lessons, err := i.inquiryService.ListLessons(ctx, &domain.ListLessonsRequest{UserId: request.GetUserId(), Page: domain.Page{
		Page: request.GetPage().GetPage(),
		Size: request.GetPage().GetSize(),
	}, Filter: domain.Filter{
		SubjectId:  request.GetFilter().GetSubjectId(),
		ConceptId:  request.GetFilter().GetConceptId(),
		UnitId:     request.GetFilter().GetUnitId(),
		SkillId:    request.GetFilter().GetSkillId(),
		GradeId:    request.GetFilter().GetGradeId(),
		SearchText: request.GetFilter().GetSearchText(),
	}})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error in ListLessons: %v", err)
	}

	resp := make([]*inquiry.Lesson, len(lessons.Lessons))

	for k, v := range lessons.Lessons {
		resp[k] = mappings.FormLesson(v)
	}
	return &inquiry.ListLessonsResponse{
		Lessons: resp,
		Count:   lessons.Count,
	}, nil
}
