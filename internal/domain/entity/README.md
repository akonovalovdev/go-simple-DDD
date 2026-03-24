### This directory contains domain entities and aggregates.

An entity is a domain object that has a unique identity that persists over time and across state changes.
An aggregate is a cluster of entities and value objects treated as a single unit with one aggregate root.

Guidelines:
- Entities should encapsulate their own invariants and expose behaviour through methods, not just getters/setters.
- Aggregates enforce consistency boundaries — all mutations go through the aggregate root.
- Do not place persistence or serialization logic here.
