# Architecture

## Overview

TaskSmith is a background job orchestration system designed to reliably schedule, execute, retry, and monitor asynchronous workloads.

The system separates **job definition**, **scheduling**, **execution**, and **operational control** into independent components. This separation keeps the architecture modular, scalable, and resilient to failures.

The core philosophy behind TaskSmith is that every component should have a single responsibility and all execution state should be persisted instead of relying on in-memory coordination.

---

## High-Level Architecture

```text
                +----------------------+
                |    Next.js UI        |
                |  (Control Plane)     |
                +----------+-----------+
                           |
                           | HTTP
                           v
                +----------------------+
                |    TaskSmith API     |
                +----------+-----------+
                           |
                           v
                +----------------------+
                |     PostgreSQL       |
                +----------+-----------+
                           ^
                           |
                   Scheduler
                           |
                           v
                    QueueForge
                           |
                           v
                    Worker Pool
```

---

## Core Components

### TaskSmith API

The API layer is responsible for accepting user requests and persisting job definitions.

Responsibilities:

- Create and manage background jobs
- Validate user input
- Persist job definitions
- Expose execution history
- Provide administrative operations

The API does **not** execute jobs directly.

---

### PostgreSQL

PostgreSQL is the system of record for TaskSmith.

It stores:

- Job definitions
- Scheduling metadata
- Execution history
- Retry information
- Execution state

All components rely on the database as the authoritative source of truth.

---

### Scheduler

The scheduler continuously evaluates persisted jobs and determines when they should execute.

Its responsibilities include:

- Finding jobs that are ready to run
- Creating JobExecution records
- Enqueuing execution requests into QueueForge

The scheduler never performs task execution.

---

### QueueForge

QueueForge provides durable message delivery between the scheduler and workers.

Its responsibilities include:

- Durable persistence
- At-least-once delivery
- Crash recovery
- Worker coordination

QueueForge is treated as an internal infrastructure dependency and is abstracted behind a queue interface.

---

### Worker Pool

Workers consume execution requests from QueueForge and execute the corresponding tasks.

Workers are intentionally stateless.

Responsibilities:

- Execute background tasks
- Report execution outcomes
- Update execution state
- Respect execution timeouts

Workers do not decide retry policies or scheduling.

---

### Control Plane

The Next.js application serves as the operational interface for TaskSmith.

Operators can:

- View jobs
- Inspect execution history
- Investigate failures
- Retry failed executions
- Pause or disable jobs

The control plane communicates only with the TaskSmith API.

---

## Design Principles

TaskSmith follows several core architectural principles.

### Separation of Intent and Execution

A Job represents **what should happen**.

A JobExecution represents **an individual attempt to execute that work**.

This distinction enables complete execution history and reliable retry behavior.

---

### Stateless Execution

Schedulers and workers maintain no persistent runtime state.

System state is always reconstructed from PostgreSQL and QueueForge.

This enables:

- Horizontal scaling
- Crash recovery
- Simplified deployments

---

### Durable State

Execution progress is never inferred from memory.

Instead, TaskSmith persists state transitions explicitly, allowing the system to recover safely after restarts.

---

### Clear Ownership

Each component owns a single responsibility.

| Component     | Responsibility            |
| ------------- | ------------------------- |
| API           | Job management            |
| PostgreSQL    | Persistent state          |
| Scheduler     | Execution planning        |
| QueueForge    | Reliable delivery         |
| Workers       | Task execution            |
| Control Plane | Operations and monitoring |

---

## Current Status

The current implementation includes:

- Project structure
- Domain models
- PostgreSQL integration
- Initial database schema
- Repository layer (in progress)

The scheduler, worker runtime, retry engine, QueueForge integration, and control plane are planned for future iterations.
