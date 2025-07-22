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
	ID              uint           `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	UserID          uuid.UUID      `json:"user_id" gorm:"type:uuid;not null" validate:"required"`
	User            User           `json:"user"`
	RoasterID       uint           `json:"roaster_id" gorm:"not null" validate:"required"`
	Roaster         Roaster        `json:"roaster"`
	CountryID       uint           `json:"country_id" gorm:"not null" validate:"required"`
	Country         Country        `json:"country"`
	AreaID          *uint          `json:"area_id"`
	Area            *Area          `json:"area" gorm:"foreignKey:AreaID"`
	FarmID          *uint          `json:"farm_id"`
	Farm            *Farm          `json:"farm" gorm:"foreignKey:FarmID"`
	FarmerID        *uint          `json:"farmer_id"`
	Farmer          *Farmer        `json:"farmer" gorm:"foreignKey:FarmerID"`
	Varieties       []Variety      `json:"variety" gorm:"many2many:bean_varieties;"`
	ProcessMethodID *uint          `json:"process_method_id"`
	ProcessMethod   *ProcessMethod `json:"process_method" gorm:"foreignKey:ProcessMethodID"`
	Name            *string        `json:"name"`
	RoastLevel      RoastLevelType `json:"roast_level" gorm:"not null;default:Medium" validate:"required"`
	Price           *uint          `json:"price" gorm:"default:null"`
	Currency        Currency       `json:"currency" gorm:"default:JPY"`
	ImageKey        *string        `json:"image_key" gorm:"default:null"`
	BeanRatings     []BeanRating   `json:"bean_ratings" gorm:"hasMany:BeanRatings;foreignKey:BeanID"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type BeanResponse struct {
	ID            uint           `json:"id" param:"id"`
	Name          *string        `json:"name,omitempty"`
	User          User           `json:"user"`
	Roaster       Roaster        `json:"roaster"`
	Country       Country        `json:"country"`
	Area          *Area          `json:"area,omitempty"`
	Farm          *Farm          `json:"farm,omitempty"`
	Farmer        *Farmer        `json:"farmer,omitempty"`
	ProcessMethod *ProcessMethod `json:"process_method,omitempty"`
	Varieties     []Variety      `json:"varieties,omitempty"`
	RoastLevel    RoastLevelType `json:"roast_level"`
	Price         *PriceResponse `json:"price,omitempty"`
	BeanRatings   []BeanRating   `json:"bean_ratings"`
	ImageURL      *string        `json:"image_url,omitempty"`
}

type PriceResponse struct {
	Amount   float64  `json:"amount"`
	Currency Currency `json:"currency"`
}

func (b *Bean) ToResponse(imageURL string) BeanResponse {
	var price *PriceResponse
	if b.Price != nil {
		price = &PriceResponse{
			Amount:   float64(*b.Price),
			Currency: b.Currency,
		}
	}

	return BeanResponse{
		ID:            b.ID,
		Name:          b.Name,
		User:          b.User,
		Roaster:       b.Roaster,
		Country:       b.Country,
		Area:          b.Area,
		Farm:          b.Farm,
		Farmer:        b.Farmer,
		ProcessMethod: b.ProcessMethod,
		Varieties:     b.Varieties,
		RoastLevel:    b.RoastLevel,
		BeanRatings:   b.BeanRatings,
		Price:         price,
		ImageURL:      &imageURL,
	}
}

type BeansResponse struct {
	Beans []BeanResponse `json:"beans"`
	Count uint           `json:"count"`
}

// roast level
type RoastLevelResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RoastLevelsResponse struct {
	RoastLevels []RoastLevelResponse `json:"roast_levels"`
	Count       uint                 `json:"count"`
}

func GetAllRoastLevels() []RoastLevelResponse {
	roastLevels := make([]RoastLevelResponse, len(AllRoastLevels))

	for i, level := range AllRoastLevels {
		roastLevels[i] = RoastLevelResponse{
			ID:   i + 1,
			Name: string(level),
		}
	}
	return roastLevels
}
