// configuring the server

package http

import (
	cfg "media_api/config"
	repo "media_api/internal/adapter/repo/persistent"
	usecase "media_api/internal/usecase"
	server "media_api/pkg/httpserver"
)

func ConfigureHttpServer(
	http_config *cfg.HTTP,
	repo *repo.RentalRepository,
	usecase *usecase.RentalUseCase,
) *server.HttpServer {
	// #TODO:
	//	add main handler
	//	add subhandlers
	//	put all subhandlers to the main one
	server := server.New(http_config)

}
