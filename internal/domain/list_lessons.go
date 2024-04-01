package domain

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (i *InquiryService) ListLessons(ctx context.Context, req *domain.ListLessonsRequest) ([]*domain.Lesson, error) {
	return i.database.ListLessons(ctx, req)
}
