// provide a usecase for
// the basic business
// needs
package usecase

import (
	"context"
	"media_api/internal/entity"
	pg_repo "media_api/internal/repo/persistent"
)

// define a structure
// of the usecase
type RentalUseCase struct {
	pg_repo *pg_repo.RentalRepository
}

// initialize a new
// usecase
func New(repo *pg_repo.RentalRepository) *RentalUseCase {
	return &RentalUseCase{repo}
}

// define the methods
// for the usecase

// Film -.
func (uc *RentalUseCase) GetFilmById(ctx context.Context, id int) (entity.Film, error) {
	film, err := uc.pg_repo.GetFilmById(ctx, id)
	if err != nil {
		return entity.Film{}, err
	}
	return film, nil
}

// Film list -.
func (uc *RentalUseCase) GetFilmsList(ctx context.Context) ([]entity.Film, error) {
	films, err := uc.pg_repo.GetFilmsList(ctx)
	if err != nil {
		return nil, err
	}
	return films, nil
}
