package postgres

import (
	"go.uber.org/zap"
	"net/url"
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

	return config, nil
}

func WithPgURL(pgURL *url.URL) func(*Config) {
	return func(config *Config) {
		config.pgURL = pgURL
	}
}
