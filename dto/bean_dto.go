package dto

import (
	"mime/multipart"
)

// Request DTOs
type CreateBeanRequest struct {
	Name          *string               `json:"name"`
	Country       CountryRef            `json:"country" validate:"required"`
	Area          *AreaRef              `json:"area,omitempty"`
	Farm          *FarmRef              `json:"farm,omitempty"`
	Farmer        *FarmerRef            `json:"farmer,omitempty"`
	Varieties     []VarietyRef          `json:"varieties" validate:"required,min=1"`
	ProcessMethod *ProcessMethodRef     `json:"process_method,omitempty"`
	RoastLevel    int                   `json:"roast_level" validate:"required,min=1,max=5"`
	Roaster       RoasterRef            `json:"roaster" validate:"required"`
	Price         *uint                 `json:"price,omitempty"`
	Currency      *string               `json:"currency,omitempty"`
	ImageFile     *multipart.FileHeader `form:"image_file,omitempty"`
	ImageURL      *string               `json:"image_url,omitempty"`
}

// Response DTOs
type BeanResponse struct {
	ID            uint                  `json:"id"`
	Name          *string               `json:"name,omitempty"`
	User          UserSummary           `json:"user"`
	Roaster       RoasterSummary        `json:"roaster"`
	Country       CountrySummary        `json:"country"`
	Area          *AreaSummary          `json:"area,omitempty"`
	Farm          *FarmSummary          `json:"farm,omitempty"`
	Farmer        *FarmerSummary        `json:"farmer,omitempty"`
	ProcessMethod *ProcessMethodSummary `json:"process_method,omitempty"`
	Varieties     []VarietySummary      `json:"varieties,omitempty"`
	RoastLevel    string                `json:"roast_level"`
	Price         *PriceResponse        `json:"price,omitempty"`
	BeanRatings   []BeanRatingSummary   `json:"bean_ratings"`
	ImageURL      *string               `json:"image_url,omitempty"`
	CreatedAt     string                `json:"created_at"`
	UpdatedAt     string                `json:"updated_at"`
}

type BeansResponse struct {
	Beans []BeanResponse `json:"beans"`
	Count uint           `json:"count"`
}

type PriceResponse struct {
	Amount         float64 `json:"amount"`
	Currency       string  `json:"currency"`
	FormattedPrice string  `json:"formatted_price"`
}

type RoastLevelResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RoastLevelsResponse struct {
	RoastLevels []RoastLevelResponse `json:"roast_levels"`
	Count       uint                 `json:"count"`
}

// Reference DTOs
type CountryRef struct {
	ID uint `json:"id" validate:"required"`
}

type AreaRef struct {
	ID uint `json:"id" validate:"required"`
}

type FarmRef struct {
	ID uint `json:"id" validate:"required"`
}

type FarmerRef struct {
	ID uint `json:"id" validate:"required"`
}

type VarietyRef struct {
	ID uint `json:"id" validate:"required"`
}

type ProcessMethodRef struct {
	ID uint `json:"id" validate:"required"`
}

type RoasterRef struct {
	ID uint `json:"id" validate:"required"`
}

// Summary DTOs for nested objects
type UserSummary struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RoasterSummary struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CountrySummary struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type AreaSummary struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type FarmSummary struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type FarmerSummary struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ProcessMethodSummary struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type VarietySummary struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type BeanRatingSummary struct {
	ID         uint    `json:"id"`
	Bitterness int     `json:"bitterness"`
	Acidity    int     `json:"acidity"`
	Body       int     `json:"body"`
	FlavorNote *string `json:"flavor_note,omitempty"`
}
