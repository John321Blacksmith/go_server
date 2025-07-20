package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	usecase "media_api/internal/usecase"
	"net/http"
)

// it should imitate the http.Handler properties
type FilmHandler struct {
	usecase *usecase.RentalUseCase
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

func (h *FilmHandler) GetFilmsList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		films, err := h.usecase.GetFilmsList(context.Background())
		if err != nil {
			w.WriteHeader(500)
			w.Write(prepareMessage(fmt.Sprintf("Something wrong has happened: %v", err)))
		}
		w.WriteHeader(200)
		w.Write(prepareMessage(films))
	})
}

func (h *FilmHandler) GetFilmById() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var pk string
		pk = r.URL.Query().Get("pk")
		film, err := h.usecase.GetFilmById(context.Background(), pk)
		if err != nil {
			w.WriteHeader(500)
			w.Write(prepareMessage(fmt.Sprintf("Some problems have occurred: %v", err)))
		}
		w.WriteHeader(200)
		w.Write(prepareMessage(film))
	})
}
