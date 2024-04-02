package mappings

import (
	"time"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
)

func ToLesson(request *inquiry.CreateLessonRequest) *domain.Lesson {
	return &domain.Lesson{
		Unit:        &domain.IdName{Id: request.GetUnitId()},
		Duration:    request.GetDuration(),
		Text:        request.GetText(),
		Name:        request.GetName(),
		UserId:      request.GetUserId(),
		Description: request.GetDescription(),
		GradeId:     request.GetGradeId(),
		Subject:     &domain.IdName{Id: request.GetSubjectId()},
		ImageId:     request.GetUserId(),
		Concept:     &domain.IdName{Id: request.GetConceptId()},
		Skill:       &domain.IdName{Id: request.GetSkillId()},
	}
}

func FormLesson(request *domain.Lesson) *inquiry.Lesson {
	return &inquiry.Lesson{
		Id: request.Id,
		Unit: &inquiry.IdName{
			Id:   request.Unit.Id,
			Name: request.Unit.Name,
		},
		Duration:    request.Duration,
		Name:        request.Name,
		Text:        request.Text,
		UserId:      request.UserId,
		Description: request.Description,
		GradeId:     request.GradeId,
		Subject: &inquiry.IdName{
			Id:   request.Subject.Id,
			Name: request.Subject.Name,
		},
		CreatedAt: request.CreatedAt.Format(time.DateTime),
		ImageId:   request.ImageId,
		Concept: &inquiry.IdName{
			Id:   request.Concept.Id,
			Name: request.Concept.Name,
		},
		Skill: &inquiry.IdName{
			Id:   request.Skill.Id,
			Name: request.Skill.Name,
		},
	}
}
