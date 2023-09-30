# monorepo-go

## Tech Stack

* Monorepo + Microservices
  * go.work
  * Earthly
* gRPC
  * connect-go
* PostgreSQL
  * pgx
  * Bun
* Passwordless Authentication

## Coding & Design Rules

## Two-Layer Rule

Limit the location of release targets (go.mod) to up to the second layer/directory.

* Good
  * services/user-service/v0.2.0  => `services/user-service@v0.2.0`
  * shared/util/v0.0.1
* Bad
  * shared/utils/sliceutil/v0.0.1  => `shared/utils/sliceutil@v0.0.1` Too long and verbose.
  * shared/helpers/dbhelper/v0.0.3
