package database

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/romanzimoglyad/inquiry-backend/internal/database/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (d *Database) ListSubjects(ctx context.Context) ([]*domain.Subject, error) {
	query, args, err := sq.Select("*").From(model.SubjectTableName.String()).PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error in selecting subjects : %w", err)
	}

	var subjects model.Subjects

	if err := pgxscan.Select(ctx, d.pool, &subjects, query, args...); err != nil {
		return nil, err
	}

	return subjects.ToDomain(), nil
}
