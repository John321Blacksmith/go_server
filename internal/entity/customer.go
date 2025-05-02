package entity

import "time"

type Customer struct {
	CustomerId int
	StoreId    int
	FirstName  string
	LastName   string
	Email      string
	AddressId  int
	ActiveBool bool
	CreateDate time.Time
	LastUpdate time.Time
	Active     bool
}
