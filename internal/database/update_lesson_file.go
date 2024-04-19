package database

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/romanzimoglyad/inquiry-backend/internal/database/model"
	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (d *Database) UpdateLessonFile(ctx context.Context, lesson *domain.Lesson) error {
	query, args, err := addFileFields(sq.Update(model.LessonTableName.String()).PlaceholderFormat(sq.Dollar), lesson).Where(sq.And{sq.And{sq.Eq{"id": lesson.Id}, sq.Eq{"user_id": lesson.UserId}}}).ToSql()

	if err != nil {
		return fmt.Errorf("error in selecting lessons : %w", err)
	}

	_, err = d.pool.Exec(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("error in updating lesson: %w", err)
	}
	return nil
}

func addFileFields(builder sq.UpdateBuilder, lesson *domain.Lesson) sq.UpdateBuilder {
	if lesson.Image != nil {
		builder = builder.Set("image_key", lesson.Image.Name)
	}

	return builder
}
