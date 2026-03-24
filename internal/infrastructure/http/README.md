### This directory contains the HTTP server setup and routing.

Responsibilities:
- Creating and configuring the HTTP server instance (timeouts, TLS, etc.).
- Registering routes and mapping them to controllers from `adapter/controller/`.
- Setting up API versioning and route grouping.
- Integrating middleware from `infrastructure/middleware/http/`.

This package should not contain any request handling logic — that belongs in `adapter/controller/`.
