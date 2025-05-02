package entity

import "time"

type Payment struct {
	PaymentId   int
	CustomerId  int
	StaffId     int
	RentalId    int
	Amount      float32
	PaymentDate time.Time
}
