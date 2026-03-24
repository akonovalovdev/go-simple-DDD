### This directory contains outbound adapter implementations — repositories.

Each subdirectory corresponds to a specific storage technology (e.g., `postgres/`, `redis/`, `mail/`).
Repository implementations fulfill the output port interfaces defined in `ports/output/` and handle all data persistence concerns.

Keep all SQL queries, ORM mappings, and storage-specific logic strictly within this layer.
The domain and use case layers must never depend on these implementations directly — only on the interfaces.
