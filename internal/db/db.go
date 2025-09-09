package db

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context) (*pgxpool.Pool, *redis.Client) {
	pgURL := os.Getenv("DATABASE_URL")
	if pgURL == "" {
		pgURL = "postgres://postgres:password@localhost:5432/store"
	}

	pool, err := pgxpool.New(ctx, pgURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return pool, rdb
}


