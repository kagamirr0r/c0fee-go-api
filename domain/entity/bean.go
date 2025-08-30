package entity

import (
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

type Bean struct {
	ID              uint
	UserID          uuid.UUID
	User            User
	RoasterID       uint
	Roaster         Roaster
	CountryID       uint
	Country         Country
	AreaID          *uint
	Area            *Area
	FarmID          *uint
	Farm            *Farm
	FarmerID        *uint
	Farmer          *Farmer
	Varieties       []Variety
	ProcessMethodID *uint
	ProcessMethod   *ProcessMethod
	Name            *string
	RoastLevelID    uint
	RoastLevel      RoastLevel
	Price           *uint
	Currency        Currency
	ImageKey        *string
	BeanRatings     []BeanRating
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
