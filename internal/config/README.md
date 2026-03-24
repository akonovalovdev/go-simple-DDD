### This directory contains the configuration loader and runtime config structs.

Responsible for reading configuration from YAML files and environment variables and exposing typed config structs to the rest of the application.

Subdirectories:
- `modules/` — one config module per subsystem (e.g., `db.go`, `log.go`, `server.go`, `smtp.go`, etc.)

The root of this package should expose a top-level `Config` struct that aggregates all module configs, and a `Load()` function that parses and validates the configuration at startup.
