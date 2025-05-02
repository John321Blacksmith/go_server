package entity

import "time"

type Country struct {
	CountryId  int
	Country    string
	LastUpdate time.Time
}
