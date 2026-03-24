### This directory contains the application use case implementations.

A use case represents a single cohesive piece of application logic that orchestrates domain objects to fulfill a business goal.

This package currently exists as the application layer. As the project grows, consider migrating to `usecase/` with
per-capability subdirectories as described in the target architecture schema.

Guidelines:
- Each use case should implement an interface from `ports/input/`.
- Use cases depend only on interfaces from `ports/output/` — never on concrete adapters.
- No HTTP, gRPC, or storage logic belongs here.