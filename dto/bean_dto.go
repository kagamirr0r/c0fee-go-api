package dto

import (
	"mime/multipart"
)

// Input DTOs
type CreateBeanInput struct {
	Data      string                `form:"data" validate:"required"`
	ImageFile *multipart.FileHeader `form:"image"`
}

type BeanInput struct {
	Name          *string        `json:"name"`
	Country       CountryRef     `json:"country" validate:"required"`
	Area          *IdRef         `json:"area,omitempty"`
	Farm          *IdRef         `json:"farm,omitempty"`
	Farmer        *IdRef         `json:"farmer,omitempty"`
	Varieties     []IdRef        `json:"varieties" validate:"required,min=1"`
	ProcessMethod *IdRef         `json:"process_method,omitempty"`
	RoastLevel    int            `json:"roast_level" validate:"required,min=1,max=5"`
	Roaster       RoasterRef     `json:"roaster" validate:"required"`
	Price         *uint          `json:"price,omitempty"`
	BeanRating    *BeanRatingRef `json:"bean_rating,omitempty"`
}

// Reference
type CountryRef struct {
	ID uint `json:"id" validate:"required"`
}

type RoasterRef struct {
	ID uint `json:"id" validate:"required"`
}

type BeanRatingRef struct {
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
	RoastLevel    string              `json:"roast_level"`
	Price         *PriceSummary       `json:"price,omitempty"`
	BeanRatings   []BeanRatingSummary `json:"bean_ratings"`
	ImageURL      *string             `json:"image_url,omitempty"`
	CreatedAt     string              `json:"created_at"`
	UpdatedAt     string              `json:"updated_at"`
}

type BeanRatingSummary struct {
	ID         uint          `json:"id"`
	User       IdNameSummary `json:"user"`
	Bitterness int           `json:"bitterness"`
	Acidity    int           `json:"acidity"`
	Body       int           `json:"body"`
	FlavorNote *string       `json:"flavor_note,omitempty"`
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

type BeansOutput struct {
	Beans      []BeanSummary `json:"beans"`
	Count      uint          `json:"count"`
	NextCursor *uint         `json:"next_cursor,omitempty"`
}

type PriceSummary struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type RoastLevelOutput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RoastLevelsOutput struct {
	RoastLevels []RoastLevelOutput `json:"roast_levels"`
	Count       uint               `json:"count"`
}
