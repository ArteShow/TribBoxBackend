package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Conn *pgxpool.Pool
	Path string
}

func Connect(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {
	connPool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		return nil, err
	}

	if err = connPool.Ping(ctx); err != nil {
		return nil, err
	}

	return connPool, nil
}
