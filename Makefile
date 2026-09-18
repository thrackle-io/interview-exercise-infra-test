.PHONY: generate db server setup teardown

# Regenerate Go code from proto/ (run after editing the .proto).
generate:
	buf generate

# Start Postgres (schema + seed applied automatically).
db:
	docker compose up -d db

server:
	cd server && DATABASE_URL=postgres://test:test@localhost:5433/test?sslmode=disable \
		go run ./cmd/server

# One-shot bootstrap.
setup: db generate
	cd server && go mod download && go mod tidy

# Stop and remove the Postgres container and its data.
teardown:
	docker compose down -v
