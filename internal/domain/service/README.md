### This directory contains domain services, policies, and factories.

Use this layer when a business operation does not naturally belong to a single entity or value object — for example, when it involves coordination between multiple aggregates or requires domain rules that span several concepts.

What belongs here:
- **Domain services** — stateless operations on domain objects (e.g., transferring ownership, matching two entities)
- **Policies** — encapsulated business rules or strategies (e.g., pricing policy, access policy)
- **Factories** — complex construction logic for aggregates or entities that goes beyond a simple constructor

Domain services must not depend on repositories, infrastructure, or use cases.
