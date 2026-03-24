### This directory contains the PostgreSQL implementation of repository interfaces.

Place all PostgreSQL-specific repository structs, query logic, and row-scanning code here.
Each repository struct should implement the corresponding output port interface from `ports/output/`.

This directory should only depend on the domain layer and the DB connection provided by `infrastructure/postgres/`.
