package entity

import "time"

type Actor struct {
	ActorId    int
	FirstName  string
	LastName   string
	LastUpdate time.Time
}
