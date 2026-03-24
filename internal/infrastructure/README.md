### This directory contains frameworks and drivers — the outermost infrastructure layer.

This layer wires together all technical concerns of the application. It depends on all other layers but nothing depends on it (except `cmd/`).

Subdirectories:
- `app/` — application bootstrap, dependency injection, and lifecycle management (start/stop)
- `http/` — HTTP server setup, routing, and gateway configuration
- `postgres/` — database connection, connection pool setup, and lifecycle
- `middleware/` — reusable middleware components (logging, recovery, auth, tracing, etc.)

Code here should be purely technical — no business logic should reside in this layer.
