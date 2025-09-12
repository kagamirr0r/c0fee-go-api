package process_method

import "time"

type Entity struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
