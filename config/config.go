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
		App      App
		HTTP     HTTP
		DataBase DataBase
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

	// DataBase -.
	DataBase struct {
		Host     string `env:"POSTGRES_HOST"`
		Port     string `env:"POSTGRES_PORT"`
		DB       string `env:"POSTGRES_DB"`
		Password string `env:"POSTGRES_PASSWORD"`
		User     string `env:"POSTGRES_USER"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error when parsing configs: %w", err)
	} else {
		// #TODO: Enable getting the configs from the json file
		cfg.App.Name = os.Getenv("APP_NAME")
		cfg.App.Version = os.Getenv("APP_VERSION")
		cfg.HTTP.Host = os.Getenv("HTTP_HOST")
		cfg.HTTP.Port = os.Getenv("HTTP_PORT")
		cfg.DataBase.Host = os.Getenv("POSTGRES_HOST")
		cfg.DataBase.Port = os.Getenv("POSTGRES_PORT")
		cfg.DataBase.DB = os.Getenv("POSTGRES_DB")
		cfg.DataBase.Password = os.Getenv("POSTGRES_PASSWORD")
		cfg.DataBase.User = os.Getenv("POSTGRES_USER")
	}
	return cfg, nil
}
