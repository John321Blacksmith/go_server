package entity

import "time"

type FilmObject struct {
	FilmId          int
	Title           string
	Description     string
	ReleaseYear     int
	LanguageId      int
	RentalDuration  time.Duration
	RentalRate      float32
	Length          int
	ReplacementCost float32
	Rating          []uint8
	LastUpdate      time.Time
	SpecialFeatures string
}

type FilmList struct {
	FilmId      int
	Title       string
	ReleaseYear int
	Length      int
	Rating      []uint8
}
