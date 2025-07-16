package http

// import (
// 	film_handler "media_api/internal/adapter/delivery/http/handler"
// 	"net/http"

// 	"golang.org/x/exp/slog"
// )

// type HttpHandler struct {
// 	fh *film_handler.FilmHandler
// }

// func NewHttpHandler(
// 	film_handler *film_handler.FilmHandler,
// ) *HttpHandler {
// 	slog.Info("Creation a new HTTP handler\n")
// 	return &HttpHandler{
// 		fh: film_handler,
// 	}
// }

// func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	h.fh.ServeHTTP(w, r)
// }
