package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// Config represents app config
type Config struct {
	DataSourceName string `envconfig:"data_source_name" required:"true"`
	DebugLog       bool   `envconfig:"debug_log"`
	GASURL         string `envconfig:"gas_url"`
}

// LoadConfig loads config
func LoadConfig() (*Config, error) {
	// do not care if .env does not exist.
	godotenv.Overload()

	c := &Config{}
	err := envconfig.Process("", c)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load config")
	}
	return c, nil
}
