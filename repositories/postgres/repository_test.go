package postgres_test

import (
	"github.com/cobbinma/example-graphql-api/repositories/postgres"
	"github.com/cobbinma/example-graphql-api/repositories/repositorytest"
	"testing"
)

func Test_Fake_Repository(t *testing.T) {
	config, err := postgres.NewConfig()
	if err != nil {
		t.Errorf("could not construct config : %s", err)
		return
	}

	repository, err := postgres.NewPostgres(config)
	if err != nil {
		t.Errorf("could not construct posgres repository : %s", err)
		return
	}

	suite := repositorytest.TestSuite(repository)
	for i := range suite {
		t.Run(suite[i].Name, suite[i].Test)
	}
}
