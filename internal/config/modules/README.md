### This directory contains per-subsystem configuration modules.

Each file or package in this directory defines a typed config struct for a specific subsystem of the application.
Examples of modules:
- `db.go` — database connection settings (host, port, pool size, timeouts)
- `server.go` — HTTP server settings (port, read/write timeouts, shutdown timeout)
- `log.go` — logging settings (level, format)
- `smtp.go` — email/SMTP settings
- `cors.go` — CORS policy settings

Keeping configs modular makes it easy to add, remove, or change settings for one subsystem without touching others.
