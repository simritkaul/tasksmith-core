# Data Model

## Overview

TaskSmith models background work using two primary entities:

- **Job** — Defines the work that should be performed.
- **JobExecution** — Represents a single attempt to execute that work.

Separating these concepts allows TaskSmith to maintain complete execution history while supporting retries, recurring schedules, and operational observability.

---

# Job

A **Job** represents a long-lived definition of background work.

It describes **what should be executed** and **how it should be executed**, but it does not represent an actual execution.

Typical examples include:

- Sending an email
- Calling an external HTTP endpoint
- Running a maintenance script
- Processing uploaded files

A single Job may produce many JobExecutions over its lifetime.

---

## Job Properties

A Job contains information such as:

- Unique identifier
- Name and description
- Task type
- Task payload
- Scheduling configuration
- Retry policy
- Timeout configuration
- Current lifecycle state

Task payloads are stored as JSON to allow different task types without requiring schema changes.

---

## Job Lifecycle

Jobs move through a simple lifecycle.

```text
CREATED
    │
    ▼
ACTIVE
    │
 ┌──┴──┐
 │     │
 ▼     ▼
PAUSED DISABLED
```

### CREATED

The job has been created but is not yet eligible for scheduling.

### ACTIVE

The scheduler evaluates the job and creates executions whenever it becomes eligible.

### PAUSED

The scheduler temporarily ignores the job.

Existing executions continue normally.

### DISABLED

The job is permanently disabled and will no longer produce new executions.

---

# JobExecution

A **JobExecution** represents one execution attempt for a Job.

Every time the scheduler determines that a Job should run, it creates a new JobExecution.

One Job may have hundreds or thousands of executions over its lifetime.

---

## JobExecution Properties

Each execution records:

- Execution identifier
- Associated Job
- Scheduled execution time
- Start time
- Finish time
- Execution state
- Attempt number
- Retry metadata
- Failure information

Unlike Jobs, JobExecutions are immutable historical records once completed.

---

## Execution State Machine

Each execution follows a well-defined lifecycle.

```text
ENQUEUED
     │
     ▼
RUNNING
  │     │
  │     ▼
  │   SUCCESS
  │
  ▼
FAILED
  │
  ├──────────────┐
  │ Retry        │
  ▼              │
ENQUEUED         │
                 │
                 ▼
               DEAD
```

### ENQUEUED

The scheduler has created the execution and placed it into QueueForge.

### RUNNING

A worker has claimed the execution and is currently processing it.

### SUCCESS

The task completed successfully.

This is a terminal state.

### FAILED

The execution encountered an error.

Depending on the retry policy, it may either be retried or transition to DEAD.

### DEAD

The execution exhausted its retry policy.

Manual intervention is required to execute the task again.

---

# Relationship

The relationship between Jobs and JobExecutions is one-to-many.

```text
Job
 │
 ├── JobExecution #1
 ├── JobExecution #2
 ├── JobExecution #3
 └── JobExecution #4
```

A Job persists for its entire lifecycle, while JobExecutions accumulate as historical records.

---

# Scheduling Model

Scheduling decisions are made against Jobs.

Execution decisions are made against JobExecutions.

This separation allows TaskSmith to support:

- One-time jobs
- Scheduled jobs
- Recurring jobs
- Retry workflows
- Execution history
- Operational dashboards

without modifying the Job definition itself.

---

# Persistence

Both entities are stored in PostgreSQL.

The database serves as the authoritative source of truth for:

- Job definitions
- Scheduling metadata
- Execution history
- Retry counters
- State transitions

Workers and schedulers remain stateless by reconstructing system state from the database whenever necessary.

---

# Design Principles

The data model follows several guiding principles:

- Separate intent from execution.
- Persist all execution state.
- Preserve complete execution history.
- Treat retries as new execution attempts rather than modifying the original job.
- Favor explicit state transitions over implicit behavior.

This model forms the foundation for the scheduler, worker pool, retry engine, and operational control plane.
