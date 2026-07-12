# TaskSmith

TaskSmith is a background job orchestration system written in Go that schedules, executes, retries, and monitors asynchronous tasks using a durable queue and a stateless worker pool.

The project is being built as an educational yet production-inspired system to explore how modern backend infrastructure handles asynchronous workloads, scheduling, failure recovery, and observability.

---

## Why TaskSmith?

Most applications need to perform work outside the request-response lifecycle, such as sending emails, generating reports, processing images, or calling external APIs.

TaskSmith provides a reliable execution pipeline for these workloads by separating **job creation**, **scheduling**, **execution**, and **retry management** into independent components.

The project focuses on correctness, durability, and clear system boundaries rather than framework-heavy abstractions.

---

## Current Features

- Background job definitions with configurable execution policies
- PostgreSQL-backed persistence layer
- Explicit Job and JobExecution domain models
- Versioned SQL schema and database migrations
- Repository pattern for persistence abstraction
- Docker-based local PostgreSQL development environment

> **Project Status:** 🚧 Under active development

---

## Planned Features

- Time-based scheduler
- QueueForge integration for durable message delivery
- Stateless worker pool
- Configurable retry and backoff policies
- Execution state machine
- Cron-based recurring jobs
- REST API for job management
- Next.js control plane for monitoring and administration
- Metrics and execution history
- Manual retry and operational controls

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

TaskSmith separates the system into distinct responsibilities:

- **API** – Accepts job definitions and management requests.
- **Scheduler** – Determines when jobs should execute.
- **QueueForge** – Provides durable, at-least-once message delivery.
- **Workers** – Execute tasks and report outcomes.
- **PostgreSQL** – Stores jobs, execution history, and system state.

---

## Technology Stack

- **Language:** Go
- **Database:** PostgreSQL
- **Queue:** QueueForge (custom durable queue)
- **Frontend:** Next.js (planned)
- **Containerization:** Docker & Docker Compose

---

## Project Structure

```text
tasksmith-core/
├── cmd/
│   └── api/
├── internal/
│   ├── db/
│   │   └── migrations/
│   ├── execution/
│   └── job/
├── pkg/
├── docker-compose.yml
├── go.mod
└── README.md
```

As development progresses, additional packages such as the scheduler, retry engine, queue abstraction, and worker runtime will be added.

---

## Getting Started

### Prerequisites

- Go 1.22+
- Docker
- Docker Compose

### Start PostgreSQL

```bash
docker compose up -d
```

### Configure Environment

Create a `.env` file:

```env
DATABASE_URL=postgres://tasksmith:tasksmith@localhost:5432/tasksmith?sslmode=disable
```

### Run the Application

```bash
go run cmd/api/main.go
```

---

## Roadmap

- [x] Project structure
- [x] PostgreSQL integration
- [x] Domain models
- [x] Initial database schema
- [ ] Repository implementation
- [ ] Scheduler
- [ ] QueueForge integration
- [ ] Worker pool
- [ ] Retry engine
- [ ] REST API
- [ ] Next.js control plane
- [ ] Metrics & observability

---

## Documentation

Detailed design documents will be added under the `docs/` directory as the project evolves, covering:

- System architecture
- Data model
- Scheduler design
- Queue integration
- Retry strategy
- Worker execution model
- API design

---

## License

This project is licensed under the MIT License.
