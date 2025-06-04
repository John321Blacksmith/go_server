package handler

import (
	"context"
	entity "media_api/internal/entity"
	usecase "media_api/internal/usecase"
	"net/http"
)

type FilmHandler struct {
	usecase *usecase.RentalUseCase
}

func New(uc *usecase.RentalUseCase) *FilmHandler {
	return &FilmHandler{
		usecase: uc,
	}
}

func (h *FilmHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) ([]entity.Film, error) {
	return h.usecase.GetFilmsList(context.Background())
}
