package domain

import (
	"context"
	"fmt"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (i *InquiryService) HandleFiles(ctx context.Context, request *domain.AddFileRequest) (string, error) {
	err := i.handleImg(ctx, request)
	if err != nil {
		return "", err
	}

	err = i.handleMaterials(ctx, request)
	if err != nil {
		return "", err
	}

	return request.LessonId, nil
}

func (i *InquiryService) handleMaterials(ctx context.Context, request *domain.AddFileRequest) error {
	if request.Files != nil {
		r := request
		for k, file := range r.Files {
			file := &domain.File{
				Name: request.LessonId + "/" + file.Name,
				Data: file.Data,
			}
			r.Files[k].Name = file.Name
			err := i.fileStorage.Upload(file)
			if err != nil {
				return fmt.Errorf("error in Upload image: %w", err)
			}
		}
		if r.Files != nil {
			err := i.database.AddMaterials(ctx, r)
			if err != nil {
				return fmt.Errorf("error in UpdateLesson (adding image): %w", err)
			}
		}
	}
	return nil
}

func (i *InquiryService) handleImg(ctx context.Context, request *domain.AddFileRequest) error {
	if request.Img != nil {
		file := &domain.File{
			Name: request.LessonId + "/" + request.Img.Name,
			Data: request.Img.Data,
		}
		err := i.fileStorage.Upload(file)
		if err != nil {
			return fmt.Errorf("error in Upload image: %w", err)
		}
		err = i.database.UpdateLesson(ctx, &domain.Lesson{
			Id:     request.LessonId,
			UserId: request.UserId,
			Image:  file,
		})
		if err != nil {
			return fmt.Errorf("error in UpdateLesson (adding image): %w", err)
		}
	}
	return nil
}
