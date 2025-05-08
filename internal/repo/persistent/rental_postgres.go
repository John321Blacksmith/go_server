// abstraction of the
// DB oriented functionality
package persistent

import (
	"context"
	"fmt"

	"media_api/internal/entity"
	"media_api/pkg/postgres"
)

// db repository
type RentalRepository struct {
	*postgres.Postgres
}

// instantiate a new
// repository powered
// by postgres driver
func NewRepository(db *postgres.Postgres) *RentalRepository {
	return &RentalRepository{db}
}

// methods for the
// repository

// Film -.
func (repo *RentalRepository) GetFilmById(ctx context.Context, id int) (entity.Film, error) {
	film := entity.Film{}
	const query = `
		SELECT
			*
		FROM
			film
		WHERE
			film_id = $1;
	`
	result := repo.Pool.QueryRow(ctx, query, id)

	err := result.Scan(
		&film.FilmId,
		&film.Title,
		&film.Description,
		&film.ReleaseYear,
		&film.LanguageId,
		&film.RentalDuration,
		&film.RentalRate,
		&film.Length,
		&film.ReplacementCost,
		&film.Rating,
		&film.LastUpdate,
		&film.SpecialFeatures,
	)
	if err != nil {
		return entity.Film{}, fmt.Errorf("error occurred while getting DB data: %w", err)
	}
	return film, nil
}

// Film list -.
func (repo *RentalRepository) GetFilmsList(ctx context.Context) ([]entity.Film, error) {
	films := []entity.Film{}
	const query = `
		SELECT
			film_id,
			title,
			release_year,
			length,
			rating
		FROM
			film
		ORDER BY
			film_id
		LIMIT $1;
	`
	result, err := repo.Pool.Query(ctx, query, 20)

	if err != nil {
		return nil, fmt.Errorf("error occurred while getting DB data: %w", err)
	}

	defer result.Close()

	for result.Next() {
		film := entity.Film{}
		result.Scan(
			&film.FilmId,
			&film.Title,
			&film.ReleaseYear,
			&film.Length,
			&film.Rating,
		)
		films = append(films, film)
	}
	return films, nil
}
