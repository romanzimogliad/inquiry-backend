package database

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/romanzimoglyad/inquiry-backend/internal/database/model"
	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (d *Database) ListDictionary(ctx context.Context, dictionaryType domain.DictionaryType) ([]*domain.IdName, error) {
	var table string
	switch dictionaryType {
	case domain.UnknownDictionaryType:
		return nil, errors.New("Unknown dictionary type")
	case domain.ConceptDictionaryType:
		table = model.ConceptTableName.String()
	case domain.SubjectDictionaryType:
		table = model.SubjectTableName.String()
	case domain.SkillDictionaryType:
		table = model.SkillTableName.String()
	case domain.UnitDictionaryType:
		table = model.UnitTableName.String()
	case domain.GradeDictionaryType:
		table = model.GradeTableName.String()
	}
	query, args, err := sq.Select("*").From(table).PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error in selecting subjects : %w", err)
	}

	var idNames model.IdNames

	if err := pgxscan.Select(ctx, d.pool, &idNames, query, args...); err != nil {
		return nil, err
	}

	return idNames.ToDomain(), nil
}
