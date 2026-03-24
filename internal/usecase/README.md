### This directory contains application use cases — business scenarios of the application.

A use case orchestrates domain objects and calls output ports to fulfill a single business scenario.
Each subdirectory groups related use cases by business capability.

Guidelines:
- Each use case struct should depend only on interfaces from `ports/input/` and `ports/output/` — never on concrete adapters.
- Use cases implement the interfaces defined in `ports/input/`.
- Use cases should not contain domain logic — delegate it to domain entities and domain services in `domain/`.
- Use cases should not know about HTTP, gRPC, or any transport protocol.
- Keep each use case focused: one struct per scenario group, one method per scenario.

Subdirectories:
- `common/` — shared application helpers reused across multiple use cases (e.g., pagination, error mapping)
