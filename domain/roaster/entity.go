package roaster

import "time"

type Entity struct {
	ID        uint
	Name      string
	Address   string
	WebURL    string
	ImageKey  *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
