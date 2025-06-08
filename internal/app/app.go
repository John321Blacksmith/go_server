// this package contains
// a function which creates
// and starts all the servers and db
// repos and usecases

package app

import (
	"fmt"
	cfg "media_api/config"
	http_cfg "media_api/internal/adapter/delivery/http"
	pg_repo "media_api/internal/adapter/repo/persistent"
	rental_usecase "media_api/internal/usecase"
	color "media_api/pkg"
	pg_driver "media_api/pkg/postgres"

	"golang.org/x/exp/slog"
)

// this function implements
// injection of such depemdencies
// as message broker, http server,
// usecases and DB repositories
func Run(cfg *cfg.Config) error {
	// DB driver startup
	slog.Info(color.Yellow + "preparing DB connection..." + color.Reset + "\n")
	slog.Info(fmt.Sprintf("DataBase configs given -: %v\n", cfg.DataBase))
	pg, err := pg_driver.NewDB(
		fmt.Sprintf(
			"host=%s port=%v user=%s password=%s dbname=%s sslmode=%s",
			cfg.DataBase.Host,
			cfg.DataBase.Port,
			cfg.DataBase.User,
			cfg.DataBase.Password,
			cfg.DataBase.DB,
			cfg.DataBase.SSL,
		),
	)
	if err != nil {
		return fmt.Errorf(color.Red+"error occurred while connecting to the DB:"+color.Reset+"%w\n", err)
	}
	defer func() {
		pg.Close()
	}()

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
	err = server.Start()
	if err != nil {
		return fmt.Errorf(color.Red+"error starting the server:"+color.Reset+"%w\n", err)
	}
	slog.Info(color.Green+"Server started on address"+color.Reset+color.Magenta+"%v"+color.Reset+"\n", cfg.HTTP.Host+":"+cfg.HTTP.Port)
	return nil
}
