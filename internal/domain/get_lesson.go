package domain

import (
	"context"
	"fmt"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

// GetLesson gets lesson
func (i *InquiryService) GetLesson(ctx context.Context, req *domain.GetLessonsRequest) (*domain.Lesson, error) {
	lesson, err := i.database.GetLesson(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error in GetLesson: %w", err)
	}
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

	return lesson, nil
}
