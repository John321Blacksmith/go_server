// configuring the server

package http

import (
	"fmt"
	cfg "media_api/config"
	handling "media_api/internal/adapter/delivery/http/handler"
	repo "media_api/internal/adapter/repo/persistent"
	usecase "media_api/internal/usecase"
	color "media_api/pkg"
	server "media_api/pkg/httpserver"
	"net/http"

	"golang.org/x/exp/slog"
)

type ConfiguredHttpServer struct {
	srv  *http.Server
	repo *repo.RentalRepository
}

func ConfigureHttpServer(
	http_config *cfg.HTTP,
	repo *repo.RentalRepository,
	usecase *usecase.RentalUseCase,
) *ConfiguredHttpServer {

	server := server.New(http_config)
	filmHandler := handling.New(usecase)
	apiHandler := NewHttpHandler(filmHandler)

	mux := http.NewServeMux()
	mux.Handle("/films", apiHandler)

	server.Handler = mux

	return &ConfiguredHttpServer{
		srv:  server,
		repo: repo,
	}
}

func (s *ConfiguredHttpServer) Start() error {
	err := s.srv.ListenAndServe()
	if err != nil {
		return err
	}
	slog.Info(fmt.Sprintf(color.Yellow+"Server is being launched on port %v..."+color.Reset, s.srv.Addr))
	return nil
}
