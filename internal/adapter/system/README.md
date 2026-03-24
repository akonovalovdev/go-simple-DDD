### This directory contains technical outbound adapters for local runtime providers.

Place here adapters that provide system-level capabilities to the application, such as:
- Clock / time provider (allows mocking time in tests)
- UUID / ID generator
- Random token generator
- Any other local runtime utility that should be abstracted behind an interface

Each adapter here should implement a corresponding interface defined in `ports/output/` so it can be easily swapped or mocked.
