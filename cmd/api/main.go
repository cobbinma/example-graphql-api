package main

import (
	"fmt"
	"github.com/cobbinma/example-graphql-api/repositories/fakerepository"
	"go.uber.org/zap"
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

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("could not construct zap logger")
		os.Exit(1)
	}
	log := logger.Sugar()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(fakerepository.NewFake())}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
