package repository

import "context"

type ActivitiesRepo interface {
	Insert(ctx context.Context, activity *Activity) error
	SelectOne(ctx context.Context) (*Activity, error)
	SelectMany(ctx context.Context, date string) (*[]Activity, error)
}
