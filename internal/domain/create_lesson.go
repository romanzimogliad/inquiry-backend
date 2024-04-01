package domain

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (i *InquiryService) CreateLesson(ctx context.Context, req *domain.Lesson) (int32, error) {
	return i.database.CreateLesson(ctx, req)
}
