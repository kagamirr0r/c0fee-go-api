package bean

import (
	"c0fee-api/domain/summary"
	"time"

	"github.com/google/uuid"
)

type Currency string

const (
	JPY Currency = "JPY"
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	KRW Currency = "KRW"
)

var AllCurrencies = []Currency{
	JPY,
	USD,
	EUR,
	GBP,
	KRW,
}

type Entity struct {
	ID              uint
	UserID          uuid.UUID
	User            summary.User
	RoasterID       uint
	Roaster         summary.Roaster
	CountryID       uint
	Country         summary.Country
	AreaID          *uint
	Area            *summary.Area
	FarmID          *uint
	Farm            *summary.Farm
	FarmerID        *uint
	Farmer          *summary.Farmer
	Varieties       []summary.Variety
	ProcessMethodID *uint
	ProcessMethod   *summary.ProcessMethod
	Name            *string
	RoastLevelID    uint
	RoastLevel      summary.RoastLevel
	Price           *uint
	Currency        Currency
	ImageKey        *string
	BeanRatings     []summary.BeanRating
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Summary = summary.Bean

// External entity references - these will be imported from their respective domains
