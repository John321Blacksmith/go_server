package entity

import "time"

type FilmCategory struct {
	FilmId     int
	CategoryId int
	LastUpdate time.Time
}
