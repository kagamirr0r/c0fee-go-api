package area

import (
	"c0fee-api/domain/summary"
	"time"
)

type Entity struct {
	ID        uint
	Name      string
	CountryID uint
	Country   summary.Country
	Farms     []summary.Farm
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Summary = summary.Area
