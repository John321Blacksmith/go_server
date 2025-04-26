package httpserver

import (
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App    *fiber.App
	notify chan error

	address         string
	prefork         bool
	readTimeout     time.Duration
	writeTimeout    time.Duration
	shutdownTimeout time.Duration
}

const (
	_defaultAddr            = ":8000"
	_defaultreadTimeout     = 5 * time.Second
	_defaultwriteTimeout    = 5 * time.Second
	_defaultshutdownTimeout = 5 * time.Second
)

func New(opts ...Option) *Server {
	s := &Server{
		App:             nil,
		notify:          make(chan error, 1), // ?
		address:         _defaultAddr,
		readTimeout:     _defaultreadTimeout,
		writeTimeout:    _defaultwriteTimeout,
		shutdownTimeout: _defaultshutdownTimeout,
	}

	app := fiber.New(
		fiber.Config{
			Prefork:      s.prefork,
			ReadTimeout:  s.readTimeout,
			WriteTimeout: s.writeTimeout,
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
		},
	)

	// apply server instance
	// attributes assignment
	for _, opt := range opts {
		opt(s)
	}
	///
	s.App = app

	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.App.Listen(s.address)
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Sutdown() error {
	return s.App.ShutdownWithTimeout(s.shutdownTimeout)
}
