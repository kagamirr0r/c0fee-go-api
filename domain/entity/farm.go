package entity

import "time"

type Farm struct {
	ID        uint
	Name      string
	AreaID    uint
	Area      Area
	Farmers   []Farmer
	CreatedAt time.Time
	UpdatedAt time.Time
}