// abstraction of the
// DB oriented functionality
package persistent

import (
	"context"
	"database/sql"
	"fmt"

	"media_api/internal/entity"

	"golang.org/x/exp/slog"
)

// db repository
type RentalRepository struct {
	db *sql.DB
}

// instantiate a new
// repository
func NewRepository(db *sql.DB) *RentalRepository {
	return &RentalRepository{db}
}

// methods for the
// repository

// Film -.
func (repo *RentalRepository) GetFilmById(ctx context.Context, id int) (entity.FilmObject, error) {
	slog.Info("Rental Reposiory: - GetFilmObject procedure is acting...")
	film := entity.FilmObject{}
	const query = `
		SELECT
			film_id,
			title,
			description,
			release_year,
			language_id,
			rental_duration,
			rental_rate,
			length,
			replacement_cost,
			rating,
			last_update,
			special_features
		FROM
			film
		WHERE
			film_id = $1;
	`
	result := repo.db.QueryRow(query, id)

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
		return entity.FilmObject{}, fmt.Errorf("error occurred while getting DB data: %w", err)
	}
	return film, nil
}

// Film list -.
func (repo *RentalRepository) GetFilmsList(ctx context.Context) ([]entity.FilmList, error) {
	slog.Info("Rental Reposiory: - GetFilmsList procedure is acting...")
	films := []entity.FilmList{}
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
	result, err := repo.db.Query(query, 20)

	if err != nil {
		return nil, fmt.Errorf("error occurred while getting DB data: %w", err)
	}

	defer result.Close()

	for result.Next() {
		film := entity.FilmList{}
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
