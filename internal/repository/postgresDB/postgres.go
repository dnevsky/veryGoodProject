package postgresDB

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresDB(url string) (db *pgxpool.Pool, err error) {
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	db, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("error while init database: %w", err)
	}

	return db, nil
}

func Close(db *pgxpool.Pool) {
	db.Close()
}
