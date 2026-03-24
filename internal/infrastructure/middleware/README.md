### This directory contains reusable middleware components.

Middleware intercepts requests/responses at the transport layer to apply cross-cutting concerns.

Subdirectories:
- `http/` — HTTP middleware (e.g., request logging, panic recovery, authentication, CORS, request ID, rate limiting)

Guidelines:
- Middleware should be stateless and composable.
- Each middleware should have a single, well-defined responsibility.
- Avoid placing business logic inside middleware — use it only for technical cross-cutting concerns.
