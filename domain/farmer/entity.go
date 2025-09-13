package farmer

import (
	"c0fee-api/domain/summary"
	"time"
)

type Entity struct {
	ID        uint
	Name      string
	FarmID    uint
	Farm      summary.Farm
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Summary = summary.Farmer
