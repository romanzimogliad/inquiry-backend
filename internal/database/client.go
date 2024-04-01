package database

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/romanzimoglyad/inquiry-backend/internal/domain"
)

type Database struct {
	pool *pgxpool.Pool
}

var _ domain.Database = &Database{}

func New(pool *pgxpool.Pool) *Database {
	return &Database{pool: pool}
}
