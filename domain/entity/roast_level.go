package entity

import "time"

type RoastLevel struct {
	ID        uint
	Name      string
	Level     int
	CreatedAt time.Time
	UpdatedAt time.Time
}