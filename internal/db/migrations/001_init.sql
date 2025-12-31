CREATE TABLE IF NOT EXISTS jobs (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    task_type TEXT NOT NULL,
    task_payload JSONB NOT NULL,

    schedule_type TEXT NOT NULL,
    scheduled_at TIMESTAMP,
    cron_expr TEXT,

    max_retries INT NOT NULL,
    backoff_strategy TEXT NOT NULL,
    timeout_ms INT NOT NULL,

    state TEXT NOT NULL,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS job_executions (
  id TEXT PRIMARY KEY,
  job_id TEXT NOT NULL REFERENCES jobs(id),

  scheduled_time TIMESTAMP NOT NULL,
  started_at TIMESTAMP,
  finished_at TIMESTAMP,

  state TEXT NOT NULL,

  attempt_number INT NOT NULL,
  max_retries INT NOT NULL,
  last_error TEXT,
  next_retry_at TIMESTAMP,

  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS uniq_job_execution_schedule 
ON job_executions(job_id, scheduled_time);