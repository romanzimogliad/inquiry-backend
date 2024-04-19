package domain

import (
	"context"
	"fmt"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (i *InquiryService) UpdateLesson(ctx context.Context, req *domain.Lesson) error {
	err := i.database.UpdateLesson(ctx, req)
	if err != nil {
		return fmt.Errorf("error in UpdateLesson: %w", err)
	}

	return nil
}
