package execution

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type PostgresRepository struct {
    db *sql.DB;
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
    return &PostgresRepository{db: db};
}

func (r *PostgresRepository) Create(ctx context.Context, e JobExecution) error {
    query := `
        INSERT INTO job_executions (
            id, job_id,
            scheduled_time,
            state,
            attempt_number, max_retries,
            created_at, updated_at
        )
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
    `

    now := time.Now().UTC()

    _, err := r.db.ExecContext(ctx, query,
        e.ID,
        e.JobID,
        e.ScheduledTime,
        e.State,
        e.AttemptNumber,
        e.MaxRetries,
        now,
        now,
    )

    return err
}

func IsDuplicate(err error) bool {
    if err == nil {
        return false
    }
    if pqErr, ok := err.(*pq.Error); ok {
        return pqErr.Code == "23505" // unique_violation
    }
    return false
}
