package http

import (
	"encoding/json"
	film_handler "media_api/internal/adapter/delivery/http/handler"
	"net/http"
)

type HttpHandler struct {
	fh *film_handler.FilmHandler
}

func NewHttpHandler(
	film_handler *film_handler.FilmHandler,
) *HttpHandler {
	return &HttpHandler{
		fh: film_handler,
	}
}

func prepareMessage(msg any) []byte {
	result, err := json.Marshal(msg)
	if err != nil {
		return nil
	} else {
		return result
	}
}

func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func (h *HttpHandler) GetFilmsList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		filmsList, err := h.fh.GetFilmsList(w, r)
		if err != nil {
			w.WriteHeader(200)
			w.Write(prepareMessage(filmsList))
		} else {
			w.WriteHeader(404)
			w.Write(prepareMessage("No Content Found"))
		}
	} else {
		w.WriteHeader(403)
		w.Write(prepareMessage("Method Not Allowed"))
	}
}
