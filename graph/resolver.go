package graph

import "github.com/cobbinma/example-graphql-api/models"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	repository models.Repository
}

func NewResolver(repository models.Repository) *Resolver {
	return &Resolver{repository: repository}
}
