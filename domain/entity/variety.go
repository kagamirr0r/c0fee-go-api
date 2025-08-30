package entity

import "time"

type Variety struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}