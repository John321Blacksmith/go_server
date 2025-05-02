package entity

import "time"

type Rental struct {
	RentalId    int
	RentalDate  time.Time
	InventoryId int
	CustomerId  int
	ReturnDate  time.Time
	StaffId     int
	LastUpdate  time.Time
}
