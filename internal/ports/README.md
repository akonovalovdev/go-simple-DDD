### This directory contains input and output contracts of the application.

Ports are the interfaces that define how the application interacts with the outside world.
They form a strict boundary between the application core and its adapters.

Subdirectories:
- `input/` — inbound port interfaces: what the application exposes to callers (use case interfaces)
- `output/` — outbound port interfaces: what the application requires from external systems (repository and service interfaces)

No implementation code should live here — only interface definitions.
