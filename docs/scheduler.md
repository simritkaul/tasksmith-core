# Scheduler

## Overview

The scheduler is responsible for determining **when** background jobs should execute.

It continuously evaluates persisted job definitions and creates executable work whenever a job becomes eligible.

The scheduler does **not** execute jobs. Its sole responsibility is to convert scheduling intent into execution requests.

---

# Responsibilities

The scheduler performs four primary tasks:

1. Scan persisted jobs.
2. Determine which jobs are eligible for execution.
3. Create a new JobExecution.
4. Enqueue the execution into QueueForge.

After the execution has been enqueued, responsibility is transferred to the worker pool.

---

# Scheduler Flow

The scheduler operates as a periodic background process.

```text
Scheduler Tick
      │
      ▼
Load Active Jobs
      │
      ▼
Evaluate Schedule
      │
      ▼
Create JobExecution
      │
      ▼
Persist Execution
      │
      ▼
Enqueue into QueueForge
```

Each iteration is independent and can safely recover after process restarts.

---

# Scheduling Decision

The scheduler evaluates only jobs that are currently active.

For each job, it determines whether the configured schedule has become due.

Typical scheduling strategies include:

- One-time execution
- Fixed interval execution
- Cron-based scheduling

The scheduler never performs task execution itself.

---

# Execution Creation

When a job becomes eligible, the scheduler creates a new JobExecution.

Each execution contains:

- Reference to the originating Job
- Scheduled execution time
- Initial execution state
- Retry configuration
- Metadata required by workers

The execution is persisted before being submitted to QueueForge.

---

# Duplicate Prevention

A scheduler may restart or multiple scheduler instances may execute concurrently.

To prevent duplicate scheduling, TaskSmith relies on PostgreSQL rather than in-memory coordination.

A unique database constraint on:

```text
(job_id, scheduled_time)
```

ensures that only one execution can exist for a given scheduling window.

This approach makes scheduling deterministic and resilient to crashes.

---

# Queue Handoff

Once the execution has been persisted, the scheduler submits the execution identifier to QueueForge.

Only the execution identifier is placed on the queue.

Workers retrieve the remaining execution metadata directly from PostgreSQL.

This keeps queue messages small while ensuring PostgreSQL remains the source of truth.

---

# Failure Handling

The scheduler is designed to be restart-safe.

If the scheduler stops unexpectedly:

- Persisted jobs remain unchanged.
- Existing executions remain durable.
- Future scheduler iterations continue from persisted state.

The scheduler maintains no in-memory scheduling state.

---

# Component Boundaries

The scheduler intentionally avoids responsibilities outside scheduling.

It does **not**:

- Execute tasks
- Apply retry policies
- Inspect worker state
- Modify completed executions
- Manage queue visibility

These responsibilities belong to other components.

---

# Scalability

The scheduler is designed to be horizontally scalable.

Multiple scheduler instances may execute concurrently because scheduling correctness is enforced by persistent database constraints rather than distributed locking.

This minimizes operational complexity while maintaining correctness.

---

# Design Principles

The scheduler follows several architectural principles:

- Keep scheduling deterministic.
- Persist execution before queue submission.
- Maintain no runtime state.
- Delegate execution to workers.
- Rely on the database for correctness.
- Keep scheduling independent from execution.

By limiting the scheduler to a single responsibility, TaskSmith remains easier to reason about, test, and scale.
