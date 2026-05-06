# simple-bank

Product name: `simple-bank`.
Repository name: `simple-bank`.
GitHub repository: [matoanbach/simple-bank](https://github.com/matoanbach/simple-bank)

`simple-bank` is a backend engineering project that simulates core banking workflows such as user registration, login, account management, and money transfers.

It was built as a hands-on learning project to practice API development, database design, transaction handling, authentication, testing, and CI workflows using Go and PostgreSQL.

This README is based on what the code in this repository actually does today: a Go REST API backed by PostgreSQL, with JWT-based authentication, session records, and transactional transfer logic.

## What It Does

This project includes:
- User registration with password hashing.
- User login with JWT token generation.
- Authenticated account creation and account listing.
- Authenticated account lookup with ownership checks.
- Money transfer logic between accounts.
- Session storage for refresh-token-style login tracking.
- PostgreSQL migrations, generated query code, and automated tests.

In plain terms, this project behaves like a small banking backend: users can sign up, log in, create accounts, and move money between accounts while the system records balances and transfer history.

## Why I Built It

I built `simple-bank` to get hands-on practice with backend and cloud-adjacent engineering topics instead of only learning them in theory.

The project helped me work through:
- designing REST API endpoints
- structuring Go application code
- writing SQL schema and queries
- generating typed database access code with `sqlc`
- handling password hashing and JWT authentication
- implementing transactional money movement safely
- running tests against PostgreSQL
- setting up GitHub Actions CI for database-backed test runs

## Skills Demonstrated

- Go backend development
- REST API design with Gin
- PostgreSQL schema design and migrations
- transactional SQL workflow design
- authentication and session concepts
- automated testing
- CI with GitHub Actions
- local developer workflow with Make targets

## Tech Stack

- Go
- Gin
- PostgreSQL
- `sqlc`
- `golang-migrate`
- `bcrypt`
- JWT
- Viper
- Testify
- GitHub Actions

## Architecture

This diagram is split into two layers so both non-technical and technical readers can understand the project quickly.

- Top half: current runtime, CI workflow, and a suggested cloud deployment shape.
- Bottom half: internal request flow, core database model, and transfer transaction design.

![simple-bank architecture and system design](images/simple-bank-architecture.png)

## Project Layout

- App entrypoint: `main.go`
- HTTP API: `api/`
- Database migrations: `db/migration/`
- SQL queries for `sqlc`: `db/query/`
- Generated DB code + store logic: `db/sqlc/`
- Config and utility helpers: `db/util/`
- Token/auth code: `token/`
- CI workflow: `.github/workflows/test.yml`
- Images used in the README: `images/`

## Run Locally

Requirements:
- Go `1.25.x`
- Docker
- PostgreSQL client tools are optional, but useful for inspection
- `migrate` CLI if you want to run migrations manually outside the Makefile workflow

Commands:

```bash
make postgres
make createdb
make migrateup
go run main.go
```

Then open the API locally at:

```text
http://localhost:8080
```

Useful commands:

```bash
make sqlc
make test
make migratedown
```

## Environment Variables

The repository includes an `app.env` file with local development values.

Key settings in that file include:
- `ENVIRONMENT`
- `DB_DRIVER`
- `DB_SOURCE`
- `MIGRATION_URL`
- `HTTP_SERVER_ADDRESS`
- `GRPC_SERVER_ADDRESS`
- `TOKEN_SYMMETRIC_KEY`
- `ACCESS_TOKEN_DURATION`
- `REFRESH_TOKEN_DURATION`

Technical note:
- The application loads config from `app.env`, but the current `main.go` still uses hardcoded database connection constants for PostgreSQL. The README setup above follows the current code path that actually runs today.

## API Summary

Current routes:

| Method | Route | Purpose |
|---|---|---|
| `POST` | `/users` | Register a new user |
| `POST` | `/users/login` | Log in and receive tokens |
| `POST` | `/accounts` | Create an account for the authenticated user |
| `GET` | `/accounts/:id` | Get a single account if it belongs to the authenticated user |
| `GET` | `/accounts` | List accounts for the authenticated user |
| `POST` | `/transfers` | Transfer money between accounts |

Protected routes require a bearer token in the `Authorization` header.

## Data Model

The database schema is centered around five main tables:

- `users`: login identity and profile information
- `accounts`: bank accounts owned by users
- `entries`: debit and credit records tied to accounts
- `transfers`: transfer records between source and destination accounts
- `sessions`: stored login sessions and refresh-token metadata

At a high level:
- a user can own one or more accounts
- a transfer creates movement between two accounts
- each transfer also creates matching accounting entries
- sessions are stored so login activity can be tracked and extended later

## Authentication and Security

The current implementation includes:
- password hashing with `bcrypt`
- JWT token creation and validation
- authentication middleware for protected routes
- account ownership checks before returning account data
- session persistence for login sessions

This means the project is not just a CRUD demo. It includes practical access control behavior for user-owned resources.

## Transaction Design

The most important backend workflow in this project is the transfer transaction.

When a transfer is created, the application:
- creates a transfer record
- creates a debit entry for the source account
- creates a credit entry for the destination account
- updates both account balances inside a database transaction

The custom store logic also updates accounts in a stable order to reduce deadlock risk during concurrent transfers.

That logic lives in `db/sqlc/store.go`.

## Testing and CI

Automated tests are currently strongest around the database and transaction layer.

The test suite covers:
- account creation, retrieval, update, deletion, and listing
- transfer transaction behavior
- concurrent transfer handling and deadlock-oriented scenarios

CI workflow:
- File: `.github/workflows/test.yml`
- Trigger: pushes to `main`
- It starts PostgreSQL, runs migrations, and executes `go test`

## Technical Notes

This section is for engineers who want a more implementation-focused view.

### HTTP Layer

The API is implemented with `Gin`.

Key files:
- `api/server.go`
- `api/user.go`
- `api/accounts.go`
- `api/transfer.go`
- `api/middleware.go`

The router separates public user/login routes from authenticated account and transfer routes.

### Database Layer

The project uses:
- SQL migrations for schema changes
- handwritten SQL queries in `db/query/`
- generated Go bindings via `sqlc`
- a custom `Store` type for transaction orchestration

This keeps SQL explicit while still giving typed Go accessors.

### Token Layer

The token package provides:
- a token maker interface
- JWT implementation
- token payload validation

This keeps authentication logic isolated from the HTTP handlers.

## What To Improve

Engineering quality:
- Align `main.go` with config-loaded database settings instead of hardcoded DB constants.
- Add API handler tests in addition to DB/store tests.
- Add route-level validation and error response consistency improvements.
- Add better structured logging.

Authentication and product behavior:
- Complete the token renewal flow.
- Expand session management behavior.
- Add stricter authorization and audit-friendly tracking.

Developer experience:
- Add a documented `.env.example` style setup.
- Standardize local ports between `app.env`, tests, and runtime.
- Add a single command for full local bootstrap.

## Summary

`simple-bank` is a backend practice project that demonstrates API development, authentication, SQL schema design, transactional money movement, testing, and CI using Go and PostgreSQL.

For non-technical readers, it shows a practical banking-style backend.
For technical readers, it shows how I approached REST design, SQL-driven data access, transaction safety, and automated verification.
