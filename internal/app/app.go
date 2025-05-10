// this package contains
// a function which creates
// and starts all the servers and db
// repos and usecases

package app

import (
	"log"
	"media_api/config"
	pg_repo "media_api/internal/repo/persistent"
	rental_usecase "media_api/internal/usecase"
	http_server "media_api/pkg/httpserver"
	pg_driver "media_api/pkg/postgres"
)

// this function implements
// injection of such depemdencies
// as message broker, http server,
// usecases and DB repositories
func Run(cfg *config.Config) {
	// DB driver startup
	pg, err := pg_driver.New(cfg)
	if err != nil {
		log.Printf("Response from DB driver: %w", err)
	}
	defer pg.Close()

	// repository startup
	rental_repo := pg_repo.NewRepository(pg)
	// usecase startup
	rental_usecase := rental_usecase.New(rental_repo)

	// message broker startup
	// message_broker := message_broker.New()

	// initialization of the http server
	server := http_server.New(&cfg.HTTP)
	server.Start()

	// http server launch
	log.Printf("Starting the server on port %v", cfg.HTTP.Port)
	if err = server.Start(); err != nil {
		log.Fatalf("Error occurred when starting the server: %v", err)
	}
}
