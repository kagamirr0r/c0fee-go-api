package entity

import "time"

type ProcessMethod struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}