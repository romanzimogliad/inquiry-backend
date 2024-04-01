package database

import (
	"context"
	"fmt"

	"github.com/romanzimoglyad/inquiry-backend/internal/database/model"

	sq "github.com/Masterminds/squirrel"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (d *Database) CreateLesson(ctx context.Context, lesson *domain.Lesson) (int32, error) {
	var lessonId int32

	query, args, err := sq.Insert(model.LessonTableName.String()).Columns("name", "user_id", "description", "grade_id", "subject_id").PlaceholderFormat(sq.Dollar).
		Values(lesson.Name, lesson.UserId, lesson.Description, lesson.GradeId, lesson.SubjectId).Suffix("RETURNING id").ToSql()
	if err != nil {
		return 0, fmt.Errorf("error in inserting lesson : %w", err)
	}
	// делаем SELECT в пару строк без циклов и сканирования)
	row := d.pool.QueryRow(ctx, query, args...)

	err = row.Scan(&lessonId)
	if err != nil {
		return 0, fmt.Errorf("error in scan result of CreateLesson: %w", err)
	}
	return lessonId, nil
}
