package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Conn *pgxpool.Pool
}

func Connect(ctx context.Context, connectionString string) (Database, error) {
	connPool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		return Database{}, err
	}

	if err = connPool.Ping(ctx); err != nil {
		connPool.Close()
		return Database{}, err
	}

	return Database{
		Conn: connPool,
	}, nil
}
