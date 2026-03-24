### This directory contains shared application helpers used across multiple use cases.

Place here utilities that are reused by several use cases but do not belong to the domain layer.

Examples of what fits here:
- Pagination helpers (building page/offset structs from request params)
- Application-level error types and error mapping utilities
- Context helpers (extracting caller identity from context)
- Transaction management helpers

Guidelines:
- Nothing here should contain business logic — that belongs in `domain/service/`.
- Do not place use case implementations here — this is a utilities package only.
- Keep it small; if a helper grows into a full abstraction, consider promoting it to its own use case or domain service.
