package infrastructure

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbPathIsEmptyError = errors.New("database path is empty")
	dbInitError        = errors.New("database init error")
)

type Database interface {
	Close()
}

type PostgreSQL struct {
	Db *pgxpool.Pool
}

func NewPostgreSqlDb(ctx context.Context, dbPath string) (*PostgreSQL, error) {
	if dbPath == "" {
		return nil, dbPathIsEmptyError
	}

	var err error
	Db, err := pgxpool.New(ctx, dbPath)
	if err != nil {
		return nil, dbInitError
	}

	return &PostgreSQL{
		Db: Db,
	}, nil
}

func (d *PostgreSQL) Close() {
	d.Db.Close()
}
