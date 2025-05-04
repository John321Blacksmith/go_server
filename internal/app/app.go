// this package contains
// a function which creates
// and starts all the servers and db
// repos and usecases

package app

import (
	"log"
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

	// initialization of the http server
	server := httpserver.New(&cfg.HTTP)

	// http server launch
	log.Printf("Starting the server on port %v", cfg.HTTP.Port)
	err := server.Start()
	if err != nil {
		log.Fatalf("Error occurred when starting the server: %v", err)
	}
}
