### This directory contains inbound port interfaces — input contracts for the application layer.

This is where you define interfaces that represent the operations your application can perform.
These interfaces are implemented by use cases in `usecase/` and called by inbound adapters in `adapter/controller/`.

Guidelines:
- One interface per use case group (e.g., `UserUseCase`, `AuthUseCase`).
- Methods should be named after application actions, not CRUD operations (e.g., `RegisterUser`, `IssueToken`).
- Input/output types for these methods should use domain DTOs from `domain/dto/` or simple plain structs.
- No implementation, no imports from adapters or infrastructure — interfaces only.