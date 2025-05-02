// this package contains
// a function which creates
// and starts all the servers and db
// repos and usecases

package app

import (
	"media_api/config"
	"media_api/pkg/httpserver"
)

// this function implements
// injection of such depemdencies
// as message broker, http server,
// usecases and DB repositories
func Run(cfg *config.Config) {
	// repository startup
	// repo := rental_repo.New()

	// usecase startup
	// usecase := rental_usecase.New()

	// message broker startup
	// message_broker := message_broker.New()

	// http server startup
	server := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
	server.Start()
}
