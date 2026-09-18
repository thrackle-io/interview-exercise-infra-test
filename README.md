# Interview Exercise — Setup Check

A minimal shell repo to confirm your local environment is ready ahead of the
interview. It verifies that:

- `buf` is installed and can generate code from `proto/`.
- The Postgres Docker container starts and is reachable.
- The Go server builds and connects to Postgres successfully.

The actual interview exercise is set up very similarly (Postgres bound to port
`5433`), so a clean run here means you're good to go.

## Prerequisites

- Docker + Docker Compose
- Go 1.25+
- `buf` CLI  (https://buf.build/docs/installation)

## Getting started

    make setup      # start DB and generate code from proto/
    make server     # connect to Postgres and print a success message

If everything is configured correctly, `make server` prints:

    ✅ Setup verified: connected to Postgres successfully.

## Teardown

    make teardown   # stop and remove the Postgres container and its data
