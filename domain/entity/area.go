package entity

import "time"

type Area struct {
	ID        uint
	Name      string
	CountryID uint
	Country   Country
	Farms     []Farm
	CreatedAt time.Time
	UpdatedAt time.Time
}