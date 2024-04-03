package domain

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (i *InquiryService) CreateLesson(ctx context.Context, req *domain.Lesson) (string, error) {
	return i.database.CreateLesson(ctx, req)
}
