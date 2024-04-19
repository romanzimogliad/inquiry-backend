package database

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/romanzimoglyad/inquiry-backend/internal/database/model"
	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (d *Database) UpdateLesson(ctx context.Context, lesson *domain.Lesson) error {
	query, args, err := addFields(sq.Update(model.LessonTableName.String()).PlaceholderFormat(sq.Dollar), lesson).Where(sq.And{sq.And{sq.Eq{"id": lesson.Id}, sq.Eq{"user_id": lesson.UserId}}}).ToSql()

	if err != nil {
		return fmt.Errorf("error in selecting lessons : %w", err)
	}

	_, err = d.pool.Exec(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("error in updating lesson: %w", err)
	}
	return nil
}

func addFields(builder sq.UpdateBuilder, lesson *domain.Lesson) sq.UpdateBuilder {
	builder = builder.Set("description", lesson.Description)
	builder = builder.Set("text", lesson.Text)
	builder = builder.Set("name", lesson.Name)
	builder = builder.Set("unit_id", lesson.Unit.Id)
	builder = builder.Set("grade_id", lesson.GradeId)
	builder = builder.Set("subject_id", lesson.Subject.Id)
	builder = builder.Set("concept_id", lesson.Concept.Id)
	builder = builder.Set("skill_id", lesson.Skill.Id)
	return builder
}
