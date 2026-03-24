### This directory contains inbound adapter implementations — controllers.

Controllers are responsible for receiving external requests (HTTP, gRPC, etc.) and translating them into use case calls.
They should not contain any business logic — only request parsing, validation, response mapping, and error handling.

Each controller implements one or more input port interfaces defined in `ports/input/`.
