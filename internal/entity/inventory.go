package entity

import "time"

type Inventory struct {
	InventoryId int
	FilmId      int
	StoreId     int
	LastUpdate  time.Time
}
