package database

import (
	"context"
	"fmt"

	"github.com/romanzimoglyad/inquiry-backend/internal/database/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

// DeleteLesson delete lesson from database
func (d *Database) DeleteLesson(ctx context.Context, request *domain.GetLessonsRequest) error {
	builder := sq.Update(model.LessonTableName.String()).PlaceholderFormat(sq.Dollar).
		Where(sq.And{sq.Eq{"id": request.Id}, sq.Eq{"active": true}}).Set("active", false)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("error in deleting lesson : %w", err)
	}

	if _, err := d.pool.Exec(ctx, query, args...); err != nil {
		return err
	}

	return nil
}
