# Changelog

All notable changes to this project will be documented in this file.

The format is based on **Keep a Changelog**, and this project follows **Semantic Versioning**.

---

## [Unreleased]

### Added

- Initial Go project structure.
- PostgreSQL integration using `database/sql`.
- Docker Compose configuration for local PostgreSQL development.
- Job domain model representing long-lived task definitions.
- JobExecution domain model representing individual execution attempts.
- Initial SQL schema for `jobs` and `job_executions` tables.
- Repository interfaces for Jobs and Job Executions.
- Initial PostgreSQL repository implementations.
- Database connection helper.
- Project documentation (`README.md`).

### Changed

- Configuration updated to read the database connection string from the `DATABASE_URL` environment variable instead of hardcoded values.

### Planned

- Complete repository implementations.
- Database migration workflow.
- Scheduler runtime.
- QueueForge integration.
- Stateless worker pool.
- Retry and backoff engine.
- REST API.
- Next.js control plane.
- Metrics and observability.

---
