package postgres

import (
	"fmt"
	"go.uber.org/zap"
	"net/url"
	"os"
	"strings"
)

const (
	POSTGRES_HOST     = "POSTGRES_HOST"
	POSTGRES_USER     = "POSTGRES_USER"
	POSTGRES_PASSWORD = "POSTGRES_PASSWORD"
	POSTGRES_DB       = "POSTGRES_DB"
	POSTGRES_SSLMODE  = "POSTGRES_SSLMODE"
)

type Config struct {
	pgURL *url.URL
	log   *zap.SugaredLogger
}

func NewConfig(log *zap.SugaredLogger, options ...func(*Config)) (*Config, error) {
	config := &Config{
		log: log,
	}

	for i := range options {
		options[i](config)
	}

	if config.pgURL == nil {
		missing := []string{}
		host, present := os.LookupEnv(POSTGRES_HOST)
		if !present {
			missing = append(missing, POSTGRES_HOST)
		}
		user, present := os.LookupEnv(POSTGRES_USER)
		if !present {
			missing = append(missing, POSTGRES_USER)
		}
		password, present := os.LookupEnv(POSTGRES_PASSWORD)
		if !present {
			missing = append(missing, POSTGRES_PASSWORD)
		}
		name, present := os.LookupEnv(POSTGRES_DB)
		if !present {
			missing = append(missing, POSTGRES_DB)
		}
		ssl, present := os.LookupEnv(POSTGRES_SSLMODE)
		if !present {
			missing = append(missing, POSTGRES_SSLMODE)
		}

		if len(missing) > 0 {
			return nil, fmt.Errorf("postgres environment variables [%s] are missing", strings.Join(missing, ", "))
		}

		pgURL := &url.URL{
			Scheme: "postgres",
			User:   url.UserPassword(user, password),
			Host:   host,
			Path:   name,
		}
		q := pgURL.Query()
		q.Add("sslmode", ssl)
		pgURL.RawQuery = q.Encode()

		config.pgURL = pgURL
	}

	return config, nil
}

func WithPgURL(pgURL *url.URL) func(*Config) {
	return func(config *Config) {
		config.pgURL = pgURL
	}
}
