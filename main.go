package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cpartogi/wgql/graph"
	"github.com/cpartogi/wgql/graph/generated"
	"github.com/joho/godotenv"
)

const defaultPort = "7000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Println(errEnv)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
