package database

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/romanzimoglyad/inquiry-backend/internal/database/model"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (d *Database) AddMaterials(ctx context.Context, request *domain.AddFileRequest) error {

	builder := sq.Insert(model.MaterialToLessonTableName.String()).Columns("material_id", "lesson_id").PlaceholderFormat(sq.Dollar)

	for _, v := range request.Files {
		builder = builder.Values(v.Name, request.LessonId)
	}

	query, args, err := builder.ToSql()
	// делаем SELECT в пару строк без циклов и сканирования)
	_, err = d.pool.Exec(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("error in add materials: %w", err)
	}
	return nil
}
