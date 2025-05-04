// this package contains
// a body of the configs
// and a function for importing
// them

package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
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
		Host string `env:"HTTP_HOST,required"`
		Port string `env:"HTTP_PORT,required"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error when parsing configs: %w", err)
	} else {
		cfg.App.Name = os.Getenv("APP_NAME")
		cfg.App.Version = os.Getenv("APP_VERSION")
		cfg.HTTP.Host = os.Getenv("HTTP_HOST")
		cfg.HTTP.Port = os.Getenv("HTTP_PORT")
	}
	return cfg, nil
}
