package database

import (
	"context"
	"fmt"

	"github.com/romanzimoglyad/inquiry-backend/internal/database/model"

	sq "github.com/Masterminds/squirrel"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (d *Database) CreateLesson(ctx context.Context, lesson *domain.Lesson) (string, error) {
	var lessonId string

	query, args, err := sq.Insert(model.LessonTableName.String()).Columns("name", "unit_id", "text", "duration", "user_id", "description", "grade_id", "subject_id", "concept_id", "skill_id").PlaceholderFormat(sq.Dollar).
		Values(lesson.Name, lesson.Unit.Id, lesson.Text, lesson.Duration, lesson.UserId, lesson.Description, lesson.GradeId, lesson.Subject.Id, lesson.Concept.Id, lesson.Skill.Id).Suffix("RETURNING id").ToSql()
	if err != nil {
		return "", fmt.Errorf("error in inserting lesson : %w", err)
	}
	// делаем SELECT в пару строк без циклов и сканирования)
	row := d.pool.QueryRow(ctx, query, args...)

	err = row.Scan(&lessonId)
	if err != nil {
		return "", fmt.Errorf("error in scan result of CreateLesson: %w", err)
	}
	return lessonId, nil
}
