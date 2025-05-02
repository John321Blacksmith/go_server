// this package contains
// a body of the configs
// and a function for importing
// them

package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type (
	// main config
	Config struct {
		App  App
		HTTP HTTP
	}

	// Application -.
	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}

	// HTTP -.
	HTTP struct {
		Port           string `env:"HTTP_PORT,required"`
		UsePreforkMode bool   `env:"HTTP_USE_PREFORK_MODE" envDefault:"false"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("error when parsing configs: %w", err)
	}
	return cfg, nil
}
