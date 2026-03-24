### This directory contains domain value objects.

A value object is an immutable domain concept defined entirely by its attributes — it has no identity.
Two value objects with the same attributes are considered equal.

Examples: Email, PhoneNumber, Money, Address, DateRange, Status.

Guidelines:
- Value objects should be immutable — prefer constructors that validate and return a ready-to-use instance.
- Embed domain validation rules directly in the constructor (e.g., return an error for an invalid email format).
- Do not reference entities or external dependencies from value objects.
