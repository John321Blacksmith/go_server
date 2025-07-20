// this package contains
// a function which creates
// and starts all the servers and db
// repos and usecases

package app

import (
	"fmt"
	cfg "media_api/config"

	// http_cfg "media_api/internal/adapter/delivery/http"
	// film_handler "media_api/internal/adapter/delivery/http/handler"
	pg_repo "media_api/internal/adapter/repo/persistent"
	rental_usecase "media_api/internal/usecase"
	color "media_api/pkg"
	pg_driver "media_api/pkg/postgres"
	"net/http"

	"golang.org/x/exp/slog"
)

// this function implements
// injection of such dependencies
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
		return fmt.Errorf(color.Red+"error occurred while connecting to the DB:"+color.Reset+"%w", err)
	}
	defer func() {
		pg.Close()
	}()

	// repository startup
	rentalRepo := pg_repo.NewRepository(pg)

	// usecase startup
	rentalUseCase := rental_usecase.New(rentalRepo)

	// create a mux object
	mux := http.NewServeMux()

	// implement a filmHandler
	// filmHandler := film_handler.New(rentalUseCase)

	// map handlers to the url patterns in MUX

	// launch the server with the MUX configured
	// http.ListenAndServe("localhost:8080", mux)

	server := http.Server{
		Addr:    cfg.HTTP.Host + ":" + cfg.HTTP.Port,
		Handler: mux,
	}
	// // configuring and initialization
	// // of the HTTP server
	// server := http_cfg.ConfigureHttpServer(
	// 	&cfg.HTTP,
	// 	rentalUseCase,
	// )

	// server launch
	err = server.ListenAndServe()

	// slog.Info(fmt.Sprintf(color.Yellow+"Server is being launched on port %v..."+color.Reset, cfg.HTTP.Host+":"+cfg.HTTP.Port))
	// if err != nil {
	// 	return fmt.Errorf(color.Red+"error starting the server:"+color.Reset+"%w", err)
	// }
	slog.Info(color.Green+"Server started on address"+color.Reset+color.Magenta+"%v"+color.Reset+"\n", cfg.HTTP.Host+":"+cfg.HTTP.Port)
	return nil
}
