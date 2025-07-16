package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	usecase "media_api/internal/usecase"
	"net/http"
)

type FilmHandlerInterface interface {
	GetFilmById(w http.ResponseWriter, r *http.Request)
	GetFilmsList(w http.ResponseWriter, r *http.Request)

	http.Handler
}

type FilmHandler struct {
	usecase *usecase.RentalUseCase

	FilmHandlerInterface
}

func prepareMessage(msg any) []byte {
	result, err := json.Marshal(msg)
	if err != nil {
		return nil
	} else {
		return result
	}
}

func New(uc *usecase.RentalUseCase) *FilmHandler {
	slog.Info("Creation of the Film Handler\n")
	return &FilmHandler{
		usecase: uc,
	}
}

func (h *FilmHandler) GetFilmsList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		filmsList, err := h.usecase.GetFilmsList(context.Background())
		if err == nil {
			w.WriteHeader(200)
			w.Write(prepareMessage(filmsList))
		} else {
			w.WriteHeader(404)
			w.Write(prepareMessage("No Content Found"))
		}
	} else {
		w.WriteHeader(405)
		w.Write(prepareMessage("Method Not Allowed"))
	}
}

func (h *FilmHandler) GetFilmById(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		film, err := h.usecase.GetFilmById(context.Background(), id)
		if err != nil {
			w.WriteHeader(404)
			w.Write(prepareMessage("No Content Found"))
		}
		w.WriteHeader(200)
		w.Write(prepareMessage(film))
	} else {
		w.WriteHeader(405)
	}
}
