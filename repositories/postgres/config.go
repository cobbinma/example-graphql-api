package postgres

type Config struct{}

func NewConfig(options ...func(*Config)) (*Config, error) {
	config := &Config{}

	for i := range options {
		options[i](config)
	}

	return config, nil
}
