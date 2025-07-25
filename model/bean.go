package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoastLevelType string

const (
	Light       RoastLevelType = "Light"
	MediumLight RoastLevelType = "Medium-Light"
	Medium      RoastLevelType = "Medium"
	MediumDark  RoastLevelType = "Medium-Dark"
	Dark        RoastLevelType = "Dark"
)

var AllRoastLevels = []RoastLevelType{
	Light,
	MediumLight,
	Medium,
	MediumDark,
	Dark,
}

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
	ID              uint      `param:"id" gorm:"primary_key;" validate:"required"`
	UserID          uuid.UUID `gorm:"type:uuid;not null" validate:"required"`
	User            User
	RoasterID       uint `gorm:"not null" validate:"required"`
	Roaster         Roaster
	CountryID       uint `gorm:"not null" validate:"required"`
	Country         Country
	AreaID          *uint
	Area            *Area `gorm:"foreignKey:AreaID"`
	FarmID          *uint
	Farm            *Farm `gorm:"foreignKey:FarmID"`
	FarmerID        *uint
	Farmer          *Farmer   `gorm:"foreignKey:FarmerID"`
	Varieties       []Variety `gorm:"many2many:bean_varieties;"`
	ProcessMethodID *uint
	ProcessMethod   *ProcessMethod `gorm:"foreignKey:ProcessMethodID"`
	Name            *string
	RoastLevel      RoastLevelType `gorm:"not null;default:Medium" validate:"required"`
	Price           *uint          `gorm:"default:null"`
	Currency        Currency       `gorm:"default:JPY"`
	ImageKey        *string        `gorm:"default:null"`
	BeanRatings     []BeanRating   `gorm:"hasMany:BeanRatings;foreignKey:BeanID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
