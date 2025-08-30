package entity

import "time"

type Country struct {
	ID        uint
	Name      string
	Code      string
	Areas     []Area
	CreatedAt time.Time
	UpdatedAt time.Time
}