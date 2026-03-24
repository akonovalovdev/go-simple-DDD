### This directory contains database migration files managed by Goose.

Each migration file represents an incremental, versioned change to the database schema.
Migrations are applied in order by the migration runner at startup or via `make migrate`.

File naming convention:
```
00001_create_users_table.sql
00002_add_email_index.sql
```

Guidelines:
- Always provide both an `-- +goose Up` and a `-- +goose Down` section in each migration file.
- Never modify an already applied migration — add a new one instead.
- Keep migrations small and focused on a single schema change.
- Avoid placing seed/fixture data here — use a dedicated seed command in `cmd/`.
