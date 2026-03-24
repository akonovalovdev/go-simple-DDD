### This directory contains the application bootstrap, dependency injection, and lifecycle management.

This is where all components are wired together into a runnable application.
Responsibilities:
- Instantiate config, logger, DB connection, and other infrastructure components.
- Construct all adapters (repositories, controllers, system providers).
- Inject dependencies into use cases and controllers.
- Register and start all long-running processes (HTTP server, background workers, etc.).
- Handle graceful shutdown on OS signals.

Think of this as the "composition root" of the application.
