package graph

import (
	"github.com/cobbinma/example-graphql-api/models"
	"time"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	repository models.Repository
	timeClient timeClient
}

func NewResolver(repository models.Repository, options ...func(*Resolver)) *Resolver {
	resolver := &Resolver{repository: repository, timeClient: realTime{}}

	for i := range options {
		options[i](resolver)
	}

	return resolver
}

func WithFixedTime(fixed time.Time) func(*Resolver) {
	return func(r *Resolver) {
		r.timeClient = fixedTime{fixed: fixed}
	}
}

type timeClient interface {
	now() time.Time
}

var _ timeClient = (*realTime)(nil)

type realTime struct{}

func (r realTime) now() time.Time {
	return time.Now()
}

var _ timeClient = (*realTime)(nil)

type fixedTime struct {
	fixed time.Time
}

func (f fixedTime) now() time.Time {
	return f.fixed
}
