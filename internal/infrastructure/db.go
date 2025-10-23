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

type Database struct {
	Db *pgxpool.Pool
}

func NewDB(dbPath string) (*Database, error) {
	if dbPath == "" {
		return nil, dbPathIsEmptyError
	}

	var err error
	Db, err := pgxpool.New(context.Background(), dbPath)
	if err != nil {
		return nil, dbInitError
	}

	return &Database{
		Db: Db,
	}, nil
}

func (d *Database) Close() {
	d.Db.Close()
}
