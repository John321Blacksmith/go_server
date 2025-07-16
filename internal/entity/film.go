package entity

import "time"

type FilmObject struct {
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

type FilmList struct {
	FilmId      int
	Title       string
	ReleaseYear int
	Length      int
	Rating      float32
}
