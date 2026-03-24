### This directory contains the domain core of the application.

This is the heart of the system. It must have zero dependencies on any framework, database, or external library.
All business rules, invariants, and domain logic live here.

Subdirectories:
- `entity/` — entities and aggregates (objects with identity and lifecycle)
- `value/` — value objects (immutable, identity-less domain concepts)
- `service/` — domain services, policies, and factories that operate over multiple entities
- `dto/` — domain-level data transfer objects used to pass data across layer boundaries

The domain layer must never import from `usecase/`, `adapter/`, `infrastructure/`, or `ports/`.
