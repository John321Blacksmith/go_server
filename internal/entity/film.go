package entity

import "time"

type Film struct {
	FilmId          int
	Title           string
	Description     string
	ReleaseYear     time.Time
	LanguageId      int
	RentalDuration  time.Duration
	RentalRate      int
	Length          int
	ReplacementCost int
	Rating          float32
	LastUpdate      time.Time
	SpecialFeatures string
}
