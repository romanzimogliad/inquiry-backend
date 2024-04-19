package database

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/romanzimoglyad/inquiry-backend/internal/database/model"
)

func (d *Database) DeleteMaterials(ctx context.Context, request []string) error {

	builder := sq.Delete(model.MaterialToLessonTableName.String()).PlaceholderFormat(sq.Dollar).Where(sq.Eq{"material_id": request})

	query, args, err := builder.ToSql()
	// делаем SELECT в пару строк без циклов и сканирования)
	_, err = d.pool.Exec(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("error in delete materials: %w", err)
	}
	return nil
}
