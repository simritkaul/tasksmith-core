# QueueForge Integration

## Overview

TaskSmith uses **QueueForge** as its internal durable message queue to reliably transfer execution requests from the scheduler to the worker pool.

Rather than implementing queueing logic directly, TaskSmith delegates message delivery guarantees to QueueForge while retaining ownership of scheduling, execution state, and retry policies.

This separation allows each component to focus on a single responsibility and keeps the overall architecture modular.

---

# Why QueueForge?

Background job systems require a reliable mechanism for transferring work between schedulers and workers.

A simple in-memory queue is insufficient because it cannot survive process crashes or machine failures.

QueueForge provides the guarantees required by TaskSmith:

- Durable message persistence
- Crash recovery
- At-least-once delivery
- Visibility timeouts
- Safe acknowledgement semantics

TaskSmith builds orchestration on top of these delivery guarantees.

---

# Architectural Boundary

TaskSmith and QueueForge have distinct responsibilities.

| Component  | Responsibility                                                    |
| ---------- | ----------------------------------------------------------------- |
| TaskSmith  | Job scheduling, execution tracking, retry policies, observability |
| QueueForge | Durable message delivery and worker coordination                  |

This separation ensures that neither component needs to understand the internal implementation of the other.

---

# Integration Flow

TaskSmith interacts with QueueForge through a simple sequence of operations.

```text
Scheduler
    │
    │ Create JobExecution
    ▼
PostgreSQL
    │
    │ Persist execution
    ▼
QueueForge
    │
    │ Deliver execution ID
    ▼
Worker
```

The scheduler never communicates directly with workers.

Workers never communicate directly with the scheduler.

QueueForge serves as the delivery layer between both components.

---

# Queue Messages

Queue messages are intentionally lightweight.

Rather than serializing the complete Job or JobExecution, TaskSmith places only the execution identifier on the queue.

Example:

```text
ExecutionID: exec_12345
```

When a worker receives this identifier, it retrieves the remaining execution metadata from PostgreSQL.

Benefits include:

- Smaller queue messages
- Reduced duplication
- Single source of truth
- Easier schema evolution

---

# Queue Abstraction

TaskSmith does not depend directly on QueueForge throughout the codebase.

Instead, it communicates through an internal queue abstraction.

Conceptually:

```text
TaskQueue
    │
    └── QueueForge Implementation
```

This approach provides several benefits:

- Loose coupling
- Easier testing
- Future extensibility
- Clear separation between business logic and infrastructure

If another queue implementation is introduced in the future, only the infrastructure layer needs to change.

---

# Delivery Semantics

QueueForge provides **at-least-once delivery**.

This means a worker may occasionally receive the same execution more than once, particularly after crashes or network failures.

TaskSmith is designed with this guarantee in mind.

Correctness is achieved by:

- Persisting execution state in PostgreSQL
- Using explicit execution state transitions
- Ensuring workers behave idempotently

The system prioritizes reliability over avoiding duplicate delivery.

---

# Failure Recovery

The combination of PostgreSQL and QueueForge enables robust recovery from failures.

Examples include:

- Scheduler restart
- Worker crash
- Process termination
- Temporary infrastructure outages

Because execution state is persisted and queue messages are durable, the system can safely resume processing without losing work.

---

# Design Principles

The QueueForge integration follows several guiding principles:

- Delegate message delivery to QueueForge.
- Keep orchestration logic inside TaskSmith.
- Persist execution state before enqueueing work.
- Minimize queue payload size.
- Communicate through abstractions instead of concrete implementations.
- Design for restartability and failure recovery.

---

# Future Enhancements

Planned improvements to the integration include:

- Dead-letter queue support
- Message prioritization
- Delayed message scheduling
- Queue metrics and monitoring
- Multiple named queues
- Rate limiting and backpressure
- Distributed worker coordination

These features will build upon the existing integration while preserving the separation between orchestration and message delivery.
