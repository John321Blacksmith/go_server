package entity

import "time"

type Staff struct {
	StaffId    int
	FirstName  string
	LastName   string
	AddressId  int
	Email      string
	StoreId    int
	Active     bool
	Username   string
	Password   string
	LastUpdate time.Time
	Picture    byte
}
