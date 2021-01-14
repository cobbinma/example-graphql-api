package main

import (
	"fmt"
	"github.com/cobbinma/example-graphql-api/repositories/postgres"
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
		return
	}

	log := logger.Sugar()
	defer func(log *zap.SugaredLogger) {
		if err := logger.Sync(); err != nil {
			log.Errorf("could not sync logger : %s", err)
		}
	}(log)

	config, err := postgres.NewConfig(log)
	if err != nil {
		log.Fatalf("could not construct postgres config : %s", err)
	}

	pg, err := postgres.NewPostgres(config)
	if err != nil {
		log.Fatalf("could not construct postgres : %s", err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(pg)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
