package domain

import (
	"context"
	"fmt"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (i *InquiryService) ListLessons(ctx context.Context, req *domain.ListLessonsRequest) (*domain.Lessons, error) {
	if req.Page.Size == 0 {
		req.Page.Size = domain.DefaultSize
	}
	if req.Page.Page == 0 {
		req.Page.Page = domain.DefaultPage
	}
	lessons, err := i.database.ListLessons(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error in ListLessons: %w", err)
	}

	for _, lesson := range lessons.Lessons {
		if lesson.Image != nil {
			url, err := i.fileStorage.GetUrl(lesson.Image.Name)
			if err != nil {
				return nil, fmt.Errorf("error in image Download: %w", err)
			}
			lesson.Image.Url = url
		}

	}
	return lessons, nil
}
