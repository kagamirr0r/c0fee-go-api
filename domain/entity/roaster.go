package entity

import "time"

type Roaster struct {
	ID        uint
	Name      string
	Address   string
	WebURL    string
	ImageKey  *string
	CreatedAt time.Time
	UpdatedAt time.Time
}