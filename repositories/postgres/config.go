package postgres

import "net/url"

type Config struct {
	pgURL *url.URL
}

func NewConfig(options ...func(*Config)) (*Config, error) {
	config := &Config{}

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
