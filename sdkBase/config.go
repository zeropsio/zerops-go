package sdkBase

type Config struct {
	Endpoint string
}

type ConfigOption func(config *Config)

func WithCustomEndpoint(endpoint string) ConfigOption {
	return func(config *Config) {
		config.Endpoint = endpoint
	}
}

func DefaultConfig(options ...ConfigOption) Config {
	c := Config{
		"https://api.app.zerops.io/",
	}
	for _, o := range options {
		o(&c)
	}
	return c
}
