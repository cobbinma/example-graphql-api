package fakerepository_test

import (
	"github.com/cobbinma/example-graphql-api/repositories/fakerepository"
	"github.com/cobbinma/example-graphql-api/repositories/repositorytest"
	"testing"
)

func Test_Repository(t *testing.T) {
	repository := fakerepository.NewFake()
	suite := repositorytest.TestSuite(repository)

	for i := range suite {
		t.Run(suite[i].Name, suite[i].Test)
	}
}
