package domain

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

// GetLesson gets lesson
func (i *InquiryService) GetLesson(ctx context.Context, req *domain.GetLessonsRequest) (*domain.Lesson, error) {
	return i.database.GetLesson(ctx, req)
}
