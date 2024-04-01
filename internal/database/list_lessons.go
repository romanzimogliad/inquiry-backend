package database

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/romanzimoglyad/inquiry-backend/internal/database/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (d *Database) ListLessons(ctx context.Context, request *domain.ListLessonsRequest) ([]*domain.Lesson, error) {
	query, args, err := sq.Select("*").From(model.LessonTableName.String()).PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"user_id": request.UserId}).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error in selecting lessons : %w", err)
	}

	var lessons model.Lessons

	if err := pgxscan.Select(ctx, d.pool, &lessons, query, args...); err != nil {
		return nil, err
	}

	return lessons.ToDomain(), nil
}
