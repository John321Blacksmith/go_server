// provide a usecase for
// the basic business
// needs
package usecase

import (
	"context"
	pg_repo "media_api/internal/adapter/repo/persistent"
	"media_api/internal/entity"
	"strconv"

	"golang.org/x/exp/slog"
)

// define a structure
// of the usecase
type RentalUseCase struct {
	pg_repo *pg_repo.RentalRepository
}

// initialize a new
// usecase
func New(repo *pg_repo.RentalRepository) *RentalUseCase {
	slog.Info("Creation of the Rental UseCase\n")
	return &RentalUseCase{repo}
}

// define the methods
// for the usecase

// Film -.
func (uc *RentalUseCase) GetFilmById(ctx context.Context, pk string) (entity.FilmObject, error) {
	slog.Info("Rental Usecase: - GetFilmObject controller is acting...")
	id, err := strconv.Atoi(pk)
	if err != nil {
		slog.Error("an error occurred converting primary key to a DB-index")
	}
	film, err := uc.pg_repo.GetFilmById(ctx, id)
	if err != nil {
		return entity.FilmObject{}, err
	}
	return film, nil
}

// Film list -.
func (uc *RentalUseCase) GetFilmsList(ctx context.Context) ([]entity.FilmList, error) {
	slog.Info("Rental Usecase: - GetFilmsList controller is acting...")
	films, err := uc.pg_repo.GetFilmsList(ctx)
	if err != nil {
		return nil, err
	}
	return films, nil
}
