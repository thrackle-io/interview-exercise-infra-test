package main

import (
	"context"
	"log"
	"os"

	"perps-exercise/internal/db"
)

func main() {
	dsn := getenv("DATABASE_URL", "postgres://test:test@localhost:5433/test?sslmode=disable")

	pool, err := db.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer pool.Close()

	log.Println("✅ Setup verified: connected to Postgres successfully.")
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
