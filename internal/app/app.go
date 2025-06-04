// this package contains
// a function which creates
// and starts all the servers and db
// repos and usecases

package app

import (
	"fmt"
	"log"
	cfg "media_api/config"
	http_cfg "media_api/internal/adapter/delivery/http"
	pg_repo "media_api/internal/adapter/repo/persistent"
	rental_usecase "media_api/internal/usecase"
	pg_driver "media_api/pkg/postgres"
)

// this function implements
// injection of such depemdencies
// as message broker, http server,
// usecases and DB repositories
func Run(cfg *cfg.Config) error {
	// DB driver startup
	pg, err := pg_driver.NewDB(
		fmt.Sprintf(
			"host=%s port=%v user=%s password=%s dbname=%s",
			cfg.DataBase.Host,
			cfg.DataBase.Port,
			cfg.DataBase.User,
			cfg.DataBase.Password,
			cfg.DataBase.DB,
		),
	)
	log.Println(cfg.DataBase)
	if err != nil {
		return fmt.Errorf("error occurred while connecting to the DB: %w", err)
	}
	defer pg.Close()

	// repository startup
	rentalRepo := pg_repo.NewRepository(pg)
	// usecase startup
	rentalUsecase := rental_usecase.New(rentalRepo)

	// configuring and initialization
	// of the HTTP server
	server := http_cfg.ConfigureHttpServer(
		&cfg.HTTP,
		rentalRepo,
		rentalUsecase,
	)

	// server launch
	if err := server.Start(); err != nil {
		return fmt.Errorf("error starting the server: %w", err)
	}
	return nil
}
