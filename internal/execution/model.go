package execution

import "time"

type JobExecution struct {
	ID string
	JobID string

	ScheduledTime time.Time
	StartedAt *time.Time
	FinishedAt *time.Time

	State string // ENQUEUED, RUNNING, SUCCESS, FAILED, DEAD

	AttemptNumber int
	MaxRetries int
	LastError *string
	NextRetryAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}