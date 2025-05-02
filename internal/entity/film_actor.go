package entity

import "time"

type FilmActor struct {
	ActorId    int
	FilmId     int
	LastUpdate time.Time
}
