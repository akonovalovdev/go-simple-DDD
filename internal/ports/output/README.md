### This directory contains outbound port interfaces — output contracts for the application layer.

This is where you define interfaces that the application core requires from external systems.
These interfaces are implemented by adapters in `adapter/repository/` and `adapter/system/`, and injected into use cases at startup.

What belongs here:
- Repository interfaces (e.g., `UserRepository`, `OrderRepository`)
- External service interfaces (e.g., `MailSender`, `PaymentGateway`)
- Technical provider interfaces (e.g., `Clock`, `IDGenerator`)

Guidelines:
- Methods should be named from the application's perspective, not the storage's (e.g., `FindByEmail`, not `SelectUserWhereEmail`).
- Return domain entities or domain DTOs — never database models or raw query results.
- No implementation, no imports from adapters or infrastructure — interfaces only.