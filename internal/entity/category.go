package entity

import "time"

type Category struct {
	CategoryId int
	Name       string
	LastUpdate time.Time
}
