// this package contains
// a function which creates
// and starts all the servers and db
// repos and usecases

package app

import (
	"media_api/config"
	"media_api/pkg/httpserver"
)

func Run(cfg *config.Config) {
	// repository startup

	// usecase startup

	// http server startup
	server := httpserver.New(httpserver.Port(cfg.HTTP.Port))
	server.Start()
}
