package database

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"

	"github.com/romanzimoglyad/inquiry-backend/internal/database/model"

	sq "github.com/Masterminds/squirrel"
)

// GetUser ...
func (d *Database) GetUser(ctx context.Context, name string) (*domain.User, error) {
	builder := sq.Select("*").From(model.UserTableName.String()).PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"name": name})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error getting user : %w", err)
	}

	var user model.User

	if err = pgxscan.Get(ctx, d.pool, &user, query, args...); err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}
