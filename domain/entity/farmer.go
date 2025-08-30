package entity

import "time"

type Farmer struct {
	ID        uint
	Name      string
	FarmID    uint
	Farm      Farm
	CreatedAt time.Time
	UpdatedAt time.Time
}