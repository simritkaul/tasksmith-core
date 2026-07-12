# API Design

## Overview

TaskSmith exposes a REST API that serves as the primary interface for creating, managing, and monitoring background jobs.

The API acts as the entry point into the system and is consumed by the TaskSmith control plane (Next.js UI) as well as any external clients.

The API is intentionally lightweight and delegates scheduling, execution, and retry logic to the internal system components.

---

# Design Goals

The API is designed around four principles:

- Manage jobs, not executions.
- Return quickly without performing background work.
- Keep execution asynchronous.
- Expose operational visibility.

Creating a job should never execute it synchronously.

---

# Architecture

```text
                Client
                   │
                   ▼
            TaskSmith API
                   │
         Persist Job Definition
                   │
                   ▼
             PostgreSQL
                   │
            Scheduler Loop
                   │
                   ▼
              QueueForge
                   │
                   ▼
                Workers
```

The API is responsible only for persisting intent.

Execution is handled entirely by the scheduler and worker runtime.

---

# Planned Endpoints

## Jobs

### Create Job

```http
POST /api/v1/jobs
```

Creates a new background job.

The request includes:

- Task type
- Task payload
- Scheduling configuration
- Retry policy
- Timeout configuration

The API validates the request and persists the Job.

---

### List Jobs

```http
GET /api/v1/jobs
```

Returns all configured jobs along with their current lifecycle state.

---

### Get Job

```http
GET /api/v1/jobs/{jobId}
```

Returns metadata for a single job, including its scheduling configuration and current status.

---

### Pause Job

```http
POST /api/v1/jobs/{jobId}/pause
```

Marks a job as paused.

Paused jobs are ignored by the scheduler until resumed.

---

### Resume Job

```http
POST /api/v1/jobs/{jobId}/resume
```

Reactivates a paused job.

---

### Disable Job

```http
POST /api/v1/jobs/{jobId}/disable
```

Prevents future executions while preserving historical data.

---

# Executions

## List Executions

```http
GET /api/v1/jobs/{jobId}/executions
```

Returns execution history for a specific job.

Each execution includes:

- State
- Attempt number
- Scheduled time
- Start time
- Finish time
- Failure reason (if applicable)

---

## Get Execution

```http
GET /api/v1/executions/{executionId}
```

Returns detailed information about a single execution attempt.

---

## Retry Execution

```http
POST /api/v1/executions/{executionId}/retry
```

Creates a new execution for a failed or dead job.

Historical execution records remain unchanged.

---

# Response Model

Successful responses return:

- Resource identifiers
- Current lifecycle state
- Relevant timestamps
- Configuration metadata

Errors return structured JSON responses containing:

- Error code
- Human-readable message
- Additional diagnostic details (when appropriate)

---

# Authentication

Authentication is planned for the control plane.

Initial development will focus on core scheduling functionality before introducing authentication and authorization.

Potential authentication mechanisms include:

- JWT
- OAuth 2.0
- Session-based authentication

---

# API Versioning

Endpoints are versioned using URL prefixes.

Example:

```text
/api/v1/jobs
```

This allows future iterations of the API without breaking existing clients.

---

# Future Enhancements

Future API capabilities may include:

- Cron expression validation
- Bulk job operations
- Job templates
- Search and filtering
- Pagination
- Webhook notifications
- Metrics endpoints
- Health and readiness endpoints

---

# Design Principles

The API follows several guiding principles:

- Persist intent rather than performing work.
- Keep request latency low.
- Separate synchronous requests from asynchronous execution.
- Expose operational visibility.
- Maintain backward compatibility through versioning.

By limiting the API to orchestration and management responsibilities, TaskSmith remains scalable, predictable, and easy to evolve.
