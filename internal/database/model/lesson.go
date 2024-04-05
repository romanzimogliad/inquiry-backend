package model

import (
	"strings"
	"time"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

type Lesson struct {
	Id          string    `db:"id"`
	Name        string    `db:"name"`
	UnitId      int32     `db:"unit_id"`
	UnitName    string    `db:"unit_name"`
	Text        string    `db:"text"`
	Duration    int32     `db:"duration"`
	UserId      int32     `db:"user_id"`
	Description string    `db:"description"`
	GradeId     int32     `db:"grade_id"`
	SubjectId   int32     `db:"subject_id"`
	SubjectName string    `db:"subject_name"`
	ImageKey    *string   `db:"image_key"`
	CreatedAt   time.Time `db:"created_at"`
	ConceptId   int32     `db:"concept_id"`
	ConceptName string    `db:"concept_name"`
	SkillId     int32     `db:"skill_id"`
	SkillName   string    `db:"skill_name"`
	Materials   *string   `db:"material_ids"`
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

	lesson := &domain.Lesson{
		Id: l.Id,
		Unit: &domain.IdName{
			Id:   l.UnitId,
			Name: l.UnitName,
		},
		Duration:    l.Duration,
		Text:        l.Text,
		Name:        l.Name,
		UserId:      l.UserId,
		Description: l.Description,
		GradeId:     l.GradeId,
		Subject: &domain.IdName{
			Id:   l.SubjectId,
			Name: l.SubjectName,
		},
		CreatedAt: l.CreatedAt,

		Concept: &domain.IdName{
			Id:   l.ConceptId,
			Name: l.ConceptName,
		},
		Skill: &domain.IdName{
			Id:   l.SkillId,
			Name: l.SkillName,
		},
	}
	if l.ImageKey != nil {
		lesson.Image = &domain.File{Name: *l.ImageKey}
	}
	if l.Materials != nil {

		materialNames := strings.Split(*l.Materials, ",")
		materials := make([]*domain.File, len(materialNames))
		for i := range materialNames {
			materials[i] = &domain.File{Name: materialNames[i]}
		}
		lesson.Files = materials
	}
	return lesson
}
