package main

import (
	"github.com/cobbinma/example-graphql-api/repositories/fakerepository"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cobbinma/example-graphql-api/graph"
	"github.com/cobbinma/example-graphql-api/graph/generated"
)

const defaultPort = "8080"

func main() {
	port, present := os.LookupEnv("PORT")
	if !present {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(fakerepository.NewFake())}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
