package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ActivityRepo struct {
	cluster *pgxpool.Pool
}

func NewActivities(cluster *pgxpool.Pool) *ActivityRepo {
	return &ActivityRepo{cluster: cluster}
}

func (r *ActivityRepo) Insert(ctx context.Context, activity *Activity) error {

	return nil
}

func (r *ActivityRepo) SelectOne(ctx context.Context, time *string) (*Activity, error) {
	return nil, nil
}

func (r *ActivityRepo) SelectMany(ctx context.Context, date string) (*[]Activity, error) {
	return nil, nil
}
