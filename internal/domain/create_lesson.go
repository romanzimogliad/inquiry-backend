package domain

import (
	"context"
	"fmt"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (i *InquiryService) CreateLesson(ctx context.Context, req *domain.Lesson) (string, error) {
	orderId, err := i.database.CreateLesson(ctx, req)
	if err != nil {
		return "", fmt.Errorf("error in CreateLesson: %w", err)
	}

	return orderId, nil
}
