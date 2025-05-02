package entity

import "time"

type Store struct {
	StoreId        int
	AddressId      int
	ManagerStaffId int
	LastUpdate     time.Time
}
