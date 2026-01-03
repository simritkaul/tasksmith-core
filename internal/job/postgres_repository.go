package job

import (
	"context"
	"database/sql"
	"time"
)

type PostgresRepository struct {
	db *sql.DB;
}

func NewPostgresRepository (db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	};
}

func (r *PostgresRepository) Create (ctx context.Context, j Job) error {
	query := `
		INSERT INTO jobs (
            id, name, description,
            task_type, task_payload,
            schedule_type, scheduled_at, cron_expr,
            max_retries, backoff_strategy, timeout_ms,
            state, created_at, updated_at
        )
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)
	`;

	now := time.Now().UTC();

	_, err := r.db.ExecContext(ctx, query, 
		j.ID,
		j.Name,
		j.Description,
		j.TaskType,
		j.TaskPayload,
		j.ScheduleType,
		j.ScheduledAt,
		j.CronExpr,
		j.MaxRetries,
		j.BackoffStrategy,
		j.TimeoutMs,
		j.State,
		now,
		now,
	);

	return err;
}

func (r *PostgresRepository) GetByID(ctx context.Context, id string) (*Job, error) {
    query := `
        SELECT
            id, name, description,
            task_type, task_payload,
            schedule_type, scheduled_at, cron_expr,
            max_retries, backoff_strategy, timeout_ms,
            state, created_at, updated_at
        FROM jobs
        WHERE id = $1
    `

    var j Job
    err := r.db.QueryRowContext(ctx, query, id).Scan(
        &j.ID,
        &j.Name,
        &j.Description,
        &j.TaskType,
        &j.TaskPayload,
        &j.ScheduleType,
        &j.ScheduledAt,
        &j.CronExpr,
        &j.MaxRetries,
        &j.BackoffStrategy,
        &j.TimeoutMs,
        &j.State,
        &j.CreatedAt,
        &j.UpdatedAt,
    )

    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }

    return &j, nil
}
