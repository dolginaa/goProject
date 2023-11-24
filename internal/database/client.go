package database

import (
	"context"

	"github.com/dolginaa/goProject/internal/configs"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDB(ctx context.Context) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(ctx, configs.ConnString)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
