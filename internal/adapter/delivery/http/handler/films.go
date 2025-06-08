package handler

import (
	"context"
	"log/slog"
	entity "media_api/internal/entity"
	usecase "media_api/internal/usecase"
	"net/http"
)

type FilmHandler struct {
	usecase *usecase.RentalUseCase
}

func New(uc *usecase.RentalUseCase) *FilmHandler {
	slog.Info("Creation of the Film Handler\n")
	return &FilmHandler{
		usecase: uc,
	}
}

func (h *FilmHandler) GetFilmsList(w http.ResponseWriter, r *http.Request) ([]entity.Film, error) {
	return h.usecase.GetFilmsList(context.Background())
}
