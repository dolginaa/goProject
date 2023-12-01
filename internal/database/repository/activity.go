package repository

import (
	"context"
	"time"

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

func (r *ActivityRepo) SelectOne(ctx context.Context) (*Activity, error) {
	row := r.cluster.QueryRow(ctx, `SELECT description, priority, deadline FROM businesses
			WHERE priority = (SELECT MAX(priority) FROM businesses)
			OR deadline::DATE = CURRENT_DATE
		ORDER BY deadline::DATE ASC
		LIMIT 1`)
	var (
		description string
		priority    int
		deadline    time.Time
	)
	err := row.Scan(&description, &priority, &deadline)
	if err != nil {
		return nil, err
	}
	return &Activity{Description: description, Priority: priority, Deadline: deadline.Format("2006-01-02 15:04:05")}, nil
}

func (r *ActivityRepo) SelectMany(ctx context.Context, date string) (*[]Activity, error) {
	rows, err := r.cluster.Query(context.Background(), "SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Activity

	for rows.Next() {
		var (
			description string
			priority    int
			deadline    time.Time
		)
		err := rows.Scan(&description, &priority, &deadline)
		if err != nil {
			return nil, err
		}
		result = append(result, Activity{Description: description, Priority: priority, Deadline: deadline.Format("2006-01-02 15:04:05")})
	}

	return &result, nil
}
