package database

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/romanzimoglyad/inquiry-backend/internal/database/model"
	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

// GetMethodToRole ...
func (d *Database) GetMethodToRole(ctx context.Context) ([]*domain.MethodToRole, error) {
	builder := sq.Select("*").From(model.MethodToRoleTableName.String()).PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error getting mtr : %w", err)
	}

	var mtr model.MethodsToRole

	if err := pgxscan.Select(ctx, d.pool, &mtr, query, args...); err != nil {
		return nil, err
	}
	return mtr.ToDomain(), nil
}
