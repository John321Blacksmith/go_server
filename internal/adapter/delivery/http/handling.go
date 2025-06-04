package http

import (
	film_handler "media_api/internal/adapter/delivery/http/handler"
	"net/http"
)

type HttpHandler struct {
	*film_handler.FilmHandler
}

func NewHttpHandler(
	film_handler *film_handler.FilmHandler,
) *HttpHandler {
	return &HttpHandler{
		film_handler,
	}
}

func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
