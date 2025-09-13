package farm

import (
	"c0fee-api/domain/summary"
	"time"
)

type Entity struct {
	ID        uint
	Name      string
	AreaID    uint
	Area      summary.Area
	Farmers   []summary.Farmer
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Summary = summary.Farm
