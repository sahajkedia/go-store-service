package main

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sahajkedia/go-store-service/internal/db"
	"github.com/sahajkedia/go-store-service/internal/gql"
)

const defaultPort = "8080"

func main() {
	ctx := context.Background()
	pgPool, redisClient := db.Connect(ctx)

	srv := handler.NewDefaultServer(gql.NewExecutableSchema(gql.Config{Resolvers: &gql.Resolver{DB: pgPool, Cache: redisClient}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
