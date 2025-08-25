package dto

import (
	"mime/multipart"
)

// Input DTOs
type BeanFormInput struct {
	Data      string                `form:"data" validate:"required"`
	ImageFile *multipart.FileHeader `form:"image"`
}

type BeanInput struct {
	ID              *uint          `json:"id,omitempty"`
	Name            *string        `json:"name"`
	CountryID       uint           `json:"country_id" validate:"required"`
	AreaID          *uint          `json:"area_id,omitempty"`
	FarmID          *uint          `json:"farm_id,omitempty"`
	FarmerID        *uint          `json:"farmer_id,omitempty"`
	VarietyIDs      []uint         `json:"variety_ids" validate:"required,min=1"`
	ProcessMethodID *uint          `json:"process_method_id,omitempty"`
	RoastLevelID    uint           `json:"roast_level_id" validate:"required"`
	RoasterID       uint           `json:"roaster_id" validate:"required"`
	Price           *uint          `json:"price,omitempty"`
	BeanRating      *BeanRatingRef `json:"bean_rating,omitempty"`
}

// Reference
type BeanRatingRef struct {
	ID         *int    `json:"id,omitempty"` // Optional ID for existing ratings, nil if creating a new rating
	Bitterness int     `json:"bitterness" validate:"required,min=1,max=5"`
	Acidity    int     `json:"acidity" validate:"required,min=1,max=5"`
	Body       int     `json:"body" validate:"required,min=1,max=5"`
	FlavorNote *string `json:"flavor_note,omitempty"`
}

// Output DTOs
type BeanOutput struct {
	ID            uint                `json:"id"`
	Name          *string             `json:"name,omitempty"`
	User          IdNameSummary       `json:"user"`
	Roaster       IdNameSummary       `json:"roaster"`
	Country       IdNameSummary       `json:"country"`
	Area          *IdNameSummary      `json:"area,omitempty"`
	Farm          *IdNameSummary      `json:"farm,omitempty"`
	Farmer        *IdNameSummary      `json:"farmer,omitempty"`
	ProcessMethod *IdNameSummary      `json:"process_method,omitempty"`
	Varieties     []IdNameSummary     `json:"varieties,omitempty"`
	RoastLevel    IdNameSummary       `json:"roast_level"`
	Price         *PriceSummary       `json:"price,omitempty"`
	BeanRatings   []BeanRatingSummary `json:"bean_ratings"`
	ImageURL      *string             `json:"image_url,omitempty"`
	CreatedAt     string              `json:"created_at"`
	UpdatedAt     string              `json:"updated_at"`
}
type BeansOutput struct {
	Beans      []BeanSummary `json:"beans"`
	Count      uint          `json:"count"`
	NextCursor *uint         `json:"next_cursor,omitempty"`
}

type BeanSummary struct {
	ID            uint            `json:"id"`
	Name          *string         `json:"name,omitempty"`
	User          IdNameSummary   `json:"user"`
	Roaster       IdNameSummary   `json:"roaster"`
	Country       IdNameSummary   `json:"country"`
	Area          *IdNameSummary  `json:"area,omitempty"`
	Farm          *IdNameSummary  `json:"farm,omitempty"`
	Farmer        *IdNameSummary  `json:"farmer,omitempty"`
	ProcessMethod *IdNameSummary  `json:"process_method,omitempty"`
	Varieties     []IdNameSummary `json:"varieties,omitempty"`
	ImageURL      *string         `json:"image_url,omitempty"`
	CreatedAt     string          `json:"created_at"`
	UpdatedAt     string          `json:"updated_at"`
}

type PriceSummary struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
