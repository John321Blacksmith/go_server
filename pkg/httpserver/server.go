package httpserver

import (
	"net/http"

	"media_api/config"
)

func New(cfg *config.HTTP) *http.Server {
	server := &http.Server{
		Addr:         cfg.Host + ":" + cfg.Port,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	return server
}
