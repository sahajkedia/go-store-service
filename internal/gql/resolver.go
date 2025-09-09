package gql

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Resolver struct {
	DB    *pgxpool.Pool
	Cache *redis.Client
}

// Stub resolvers for demo
func (r *Resolver) CreateStore(ctx context.Context, name string) (string, error) {
	_, err := r.DB.Exec(ctx, "INSERT INTO stores(name) VALUES($1)", name)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Store %s created", name), nil
}
