package domain

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

type InquiryService struct {
	database    Database
	fileStorage FileStorage
}

type Database interface {
	Ping(ctx context.Context) (string, error)
	CreateLesson(ctx context.Context, lesson *domain.Lesson) (string, error)
	UpdateLesson(ctx context.Context, lesson *domain.Lesson) error
	ListLessons(ctx context.Context, lesson *domain.ListLessonsRequest) (*domain.Lessons, error)
	GetLesson(ctx context.Context, lesson *domain.GetLessonsRequest) (*domain.Lesson, error)
	DeleteLesson(ctx context.Context, lesson *domain.GetLessonsRequest) error
	ListSubjects(ctx context.Context) ([]*domain.Subject, error)
	ListDictionary(ctx context.Context, dictionaryType domain.DictionaryType) ([]*domain.IdName, error)
	UpdateLessonFile(ctx context.Context, lesson *domain.Lesson) error
	AddMaterials(ctx context.Context, request *domain.AddFileRequest) error
	DeleteMaterials(ctx context.Context, request []string) error
}

type FileStorage interface {
	Upload(file *domain.File) error
	Download(key string) (*domain.File, error)
	GetUrl(key string) (string, error)
}

func New(database Database, fileStorage FileStorage) *InquiryService {
	return &InquiryService{
		database:    database,
		fileStorage: fileStorage,
	}
}
