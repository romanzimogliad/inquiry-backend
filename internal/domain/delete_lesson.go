package domain

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

// DeleteLesson delete lesson
func (i *InquiryService) DeleteLesson(ctx context.Context, req *domain.GetLessonsRequest) error {
	return i.database.DeleteLesson(ctx, req)
}
