package httpserver

import (
	"net/http"
	"time"

	"media_api/config"
)

type HttpServer struct {
	implementation *http.Server
	address        string
	readTimeout    time.Duration
	writeTimeout   time.Duration
}

const (
	_defaultAddr         = ":8000"
	_defaultreadTimeout  = 5 * time.Second
	_defaultwriteTimeout = 5 * time.Second
)

func New(cfg *config.HTTP) *HttpServer {
	server := &HttpServer{
		address:      cfg.Host + ":" + cfg.Port,
		readTimeout:  _defaultreadTimeout,
		writeTimeout: _defaultwriteTimeout,
	}

	app := &http.Server{
		Addr:         server.address,
		ReadTimeout:  server.readTimeout,
		WriteTimeout: server.writeTimeout,
	}

	server.implementation = app

	return server
}

func (s *HttpServer) Start() error {
	err := s.implementation.ListenAndServe()
	if err != nil {
		return err
	} else {
		return nil
	}
}
