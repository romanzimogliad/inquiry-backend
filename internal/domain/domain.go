package domain

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

type InquiryService struct {
	database Database
}

type Database interface {
	Ping(ctx context.Context) (string, error)
	CreateLesson(ctx context.Context, lesson *domain.Lesson) (int32, error)
	ListLessons(ctx context.Context, lesson *domain.ListLessonsRequest) ([]*domain.Lesson, error)
}

func New(database Database) *InquiryService {
	return &InquiryService{
		database: database,
	}
}
