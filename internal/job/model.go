package job

import "time"

type Job struct {
	ID string
	Name string
	Description string
	TaskType string
	TaskPayload []byte

	ScheduleType string
	ScheduledAt *time.Time
	CronExpr *string

	MaxRetries int
	BackoffStrategy string
	TimeoutMs int

	State string	// CREATED, ACTIVE, PAUSED, DISABLED

	CreatedAt time.Time
	UpdatedAt time.Time
}