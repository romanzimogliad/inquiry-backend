package model

import (
	"time"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

type Lesson struct {
	Id          int32     `db:"id"`
	Name        string    `db:"name"`
	UserId      int64     `db:"user_id"`
	Description string    `db:"description"`
	GradeId     int32     `db:"grade_id"`
	SubjectId   int32     `db:"subject_id"`
	CreatedAt   time.Time `db:"created_at"`
}

type Lessons []*Lesson

func (l *Lessons) ToDomain() []*domain.Lesson {
	resp := make([]*domain.Lesson, len(*l))
	for k, v := range *l {
		resp[k] = v.ToDomain()
	}
	return resp
}

func (l *Lesson) ToDomain() *domain.Lesson {
	return &domain.Lesson{
		Id:          l.Id,
		Name:        l.Name,
		UserId:      l.UserId,
		Description: l.Description,
		GradeId:     l.GradeId,
		SubjectId:   l.SubjectId,
		CreatedAt:   l.CreatedAt,
	}
}
