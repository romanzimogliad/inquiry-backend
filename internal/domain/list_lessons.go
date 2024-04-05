package domain

import (
	"context"
	"fmt"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (i *InquiryService) ListLessons(ctx context.Context, req *domain.ListLessonsRequest) ([]*domain.Lesson, error) {
	lessons, err := i.database.ListLessons(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error in ListLessons: %w", err)
	}

	for _, lesson := range lessons {
		if lesson.Image != nil {
			url, err := i.fileStorage.GetUrl(lesson.Image.Name)
			if err != nil {
				return nil, fmt.Errorf("error in image Download: %w", err)
			}
			lesson.Image.Url = url
		}
		if lesson.Files != nil {
			for k, file := range lesson.Files {
				url, err := i.fileStorage.GetUrl(file.Name)
				if err != nil {
					return nil, fmt.Errorf("error in image Download: %w", err)
				}
				lesson.Files[k].Url = url

			}
		}
	}
	return lessons, nil
}
