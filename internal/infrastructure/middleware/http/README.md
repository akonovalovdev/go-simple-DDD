### This directory contains HTTP-specific middleware.

Place here all middleware functions that wrap HTTP handlers to provide cross-cutting concerns:
- **Logging** — log incoming requests and outgoing responses with status codes and durations
- **Recovery** — catch panics and return a 500 response instead of crashing the server
- **Authentication** — validate tokens/sessions and populate request context with identity
- **CORS** — handle cross-origin resource sharing headers
- **Request ID** — attach a unique trace ID to every request for observability
- **Rate limiting** — protect endpoints from abuse

Each middleware should follow the standard `func(http.Handler) http.Handler` signature for easy chaining.
