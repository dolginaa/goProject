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
	commTag, err := r.cluster.Exec(ctx, `INSERT INTO businesses(datetime, priority, deadline, description) VALUES($1, $2, $3, $4)`, activity.DateTime, activity.Priority, activity.Deadline, activity.Description)
	if err != nil {
		return err
	}
	//поправить ошибку при недобавлении
	if commTag.RowsAffected() == 0 {
		return err
	}
	return nil
}

func (r *ActivityRepo) SelectOne(ctx context.Context, time *string) (*Activity, error) {
	return nil, nil
}

func (r *ActivityRepo) SelectMany(ctx context.Context, date string) (*[]Activity, error) {
	return nil, nil
}
