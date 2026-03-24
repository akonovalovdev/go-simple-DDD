### This directory contains the main package of the service.

This is the entrypoint binary for the primary service.
It is responsible for wiring together all dependencies (config, infrastructure, adapters, use cases) and starting the application.
Keep this package thin — it should only bootstrap and delegate to `internal/infrastructure/app`.
