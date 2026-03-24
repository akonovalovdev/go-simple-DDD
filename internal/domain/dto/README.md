### This directory contains domain-level data transfer objects (DTOs).

Domain DTOs are plain structs used to carry data between the domain layer and the use case / application layer.
They decouple the domain model from the outer layers, allowing domain entities to evolve independently.

Guidelines:
- DTOs should be simple structs with no behaviour.
- Use them as input/output types for domain service methods and use case boundaries.
- Do not embed infrastructure or framework-specific types here.
