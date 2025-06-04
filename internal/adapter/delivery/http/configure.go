// configuring the server

package http

import (
	cfg "media_api/config"
	handling "media_api/internal/adapter/delivery/http/handler"
	repo "media_api/internal/adapter/repo/persistent"
	usecase "media_api/internal/usecase"
	server "media_api/pkg/httpserver"
	"net/http"
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
	// #TODO:
	//	add main handler
	//	add subhandlers
	//	put all subhandlers to the main one

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
	if err := s.srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
