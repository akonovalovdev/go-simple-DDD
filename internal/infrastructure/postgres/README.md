### This directory contains the PostgreSQL connection and connection pool setup.

Responsibilities:
- Opening and configuring the database connection using settings from `config/modules/`.
- Setting connection pool parameters (max open/idle connections, connection lifetime).
- Providing a shared `*sql.DB` or equivalent pool instance to repository adapters.
- Handling connection health checks and graceful close on shutdown.

No SQL queries or business logic should be placed here — this is purely a driver/infrastructure concern.
