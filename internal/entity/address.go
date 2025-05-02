package entity

import "time"

type Address struct {
	AddressId  int
	Address    string
	Address2   string
	District   string
	CityId     int
	PostalCode string
	Phone      string
	LastUpdate time.Time
}
