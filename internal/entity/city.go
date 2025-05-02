package entity

import "time"

type City struct {
	CityId     int
	City       string
	CountryId  int
	LastUpdate time.Time
}
