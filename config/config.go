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
		HTTP HTTP
	}

	// HTTP -.
	HTTP struct {
		Port string
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
