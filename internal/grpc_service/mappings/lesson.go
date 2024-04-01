package mappings

import (
	"time"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
)

func ToLesson(request *inquiry.CreateLessonRequest) *domain.Lesson {
	return &domain.Lesson{

		Name:        request.GetName(),
		UserId:      request.GetUserId(),
		Description: request.GetDescription(),
		GradeId:     request.GetGradeId(),
		SubjectId:   request.GetSubjectId(),
	}
}

func FormLesson(request *domain.Lesson) *inquiry.Lesson {
	return &inquiry.Lesson{
		Id:          request.Id,
		Name:        request.Name,
		UserId:      request.UserId,
		Description: request.Description,
		GradeId:     request.GradeId,
		SubjectId:   request.SubjectId,
		CreatedAt:   request.CreatedAt.Format(time.DateTime),
	}
}
