### This directory contains all input and output adapters.

Adapters are the glue between the outside world and the application core.
They implement the interfaces defined in `ports/` and translate between external representations and domain/application models.

Subdirectories:
- `controller/` — inbound adapters: HTTP or gRPC handlers that receive requests and invoke use cases
- `repository/` — outbound adapters: implementations of output port interfaces (e.g., database repositories)
- `system/` — technical outbound adapters: local runtime providers such as clock, UUID generator, etc.
