### This directory contains all internal application code.

Nothing inside this directory is intended to be imported by external projects.
It is organized into clearly separated layers following Clean Architecture and DDD principles:

- `adapter/` — input/output adapters (controllers, repositories, technical providers)
- `config/` — config loader and per-subsystem configuration modules
- `domain/` — the domain core (entities, value objects, domain services, DTOs)
- `infrastructure/` — frameworks and drivers (HTTP server, DB connection, DI, middleware)
- `ports/` — input and output contracts (interfaces)
- `usecase/` — application use cases (business scenarios)
