package entity

import "time"

type Language struct {
	LanguageId int
	Name       string
	LastUpdate time.Time
}
