package country

import (
	"c0fee-api/domain/summary"
	"time"
)

type Entity struct {
	ID        uint
	Name      string
	Code      string
	Areas     []summary.Area
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Summary = summary.Country
