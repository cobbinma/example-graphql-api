package postgres_test

import (
	"github.com/cobbinma/example-graphql-api/models"
	"github.com/cobbinma/example-graphql-api/repositories/postgres"
	"github.com/cobbinma/example-graphql-api/repositories/repositorytest"
	"github.com/ory/dockertest/v3"
	"net"
	"net/url"
	"runtime"
	"testing"
	"time"
)

func Test_Postgres_Repository(t *testing.T) {
	pgURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("myuser", "mypass"),
		Path:   "mydatabase",
	}
	q := pgURL.Query()
	q.Add("sslmode", "disable")
	pgURL.RawQuery = q.Encode()

	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Errorf("could not connect to docker : %s", err)
		return
	}

	pw, _ := pgURL.User.Password()
	runOpts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			"POSTGRES_USER=" + pgURL.User.Username(),
			"POSTGRES_PASSWORD=" + pw,
			"POSTGRES_DB=" + pgURL.Path,
		},
	}

	resource, err := pool.RunWithOptions(&runOpts)
	if err != nil {
		t.Errorf("Could start postgres container : %s", err)
		return
	}
	defer func() {
		err = pool.Purge(resource)
		if err != nil {
			t.Errorf("Could not purge resource : %s", err)
		}
	}()

	pgURL.Host = resource.Container.NetworkSettings.IPAddress

	// Docker layer network is different on Mac
	if runtime.GOOS == "darwin" {
		pgURL.Host = net.JoinHostPort(resource.GetBoundIP("5432/tcp"), resource.GetPort("5432/tcp"))
	}

	config, err := postgres.NewConfig(postgres.WithPgURL(pgURL))
	if err != nil {
		t.Errorf("could not construct config : %s", err)
		return
	}

	var repository models.Repository

	pool.MaxWait = 10 * time.Second
	err = pool.Retry(func() error {
		repository, err = postgres.NewPostgres(config)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		t.Errorf("could not connect to postgres server : %s", err)
		return
	}

	suite := repositorytest.TestSuite(repository)
	for i := range suite {
		t.Run(suite[i].Name, suite[i].Test)
	}
}
