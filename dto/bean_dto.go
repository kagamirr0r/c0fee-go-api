package dto

import (
	"mime/multipart"
)

// Input DTOs
type CreateBeanInput struct {
	Data      string                `form:"data" validate:"required"`
	ImageFile *multipart.FileHeader `form:"image"`
}

type CreateBeanData struct {
	Name          *string           `json:"name"`
	Country       CountryRef        `json:"country" validate:"required"`
	Area          *IdRef            `json:"area,omitempty"`
	Farm          *IdRef            `json:"farm,omitempty"`
	Farmer        *IdRef            `json:"farmer,omitempty"`
	Varieties     []IdRef           `json:"varieties" validate:"required,min=1"`
	ProcessMethod *IdRef            `json:"process_method,omitempty"`
	RoastLevel    int               `json:"roast_level" validate:"required,min=1,max=5"`
	Roaster       RoasterRef        `json:"roaster" validate:"required"`
	Price         *uint             `json:"price,omitempty"`
	BeanRating    *CreateBeanRating `json:"bean_rating,omitempty"`
}

type CreateBeanRating struct {
	Bitterness int     `json:"bitterness" validate:"required,min=1,max=5"`
	Acidity    int     `json:"acidity" validate:"required,min=1,max=5"`
	Body       int     `json:"body" validate:"required,min=1,max=5"`
	FlavorNote *string `json:"flavor_note,omitempty"`
}

type BeanOutput struct {
	ID            uint                `json:"id"`
	Name          *string             `json:"name,omitempty"`
	User          IdNameOutput        `json:"user"`
	Roaster       IdNameOutput        `json:"roaster"`
	Country       IdNameOutput        `json:"country"`
	Area          *IdNameOutput       `json:"area,omitempty"`
	Farm          *IdNameOutput       `json:"farm,omitempty"`
	Farmer        *IdNameOutput       `json:"farmer,omitempty"`
	ProcessMethod *IdNameOutput       `json:"process_method,omitempty"`
	Varieties     []IdNameOutput      `json:"varieties,omitempty"`
	RoastLevel    string              `json:"roast_level"`
	Price         *PriceOutput        `json:"price,omitempty"`
	BeanRatings   []BeanRatingSummary `json:"bean_ratings"`
	ImageURL      *string             `json:"image_url,omitempty"`
	CreatedAt     string              `json:"created_at"`
	UpdatedAt     string              `json:"updated_at"`
}

type CreateBeanOutput struct {
	Bean    BeanOutput `json:"bean"`
	Message string     `json:"message"`
}

type BeansOutput struct {
	Beans []BeanOutput `json:"beans"`
	Count uint         `json:"count"`
}

type PriceOutput struct {
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

// Reference DTOs
type CountryRef struct {
	ID uint `json:"id" validate:"required"`
}

type RoasterRef struct {
	ID uint `json:"id" validate:"required"`
}

type BeanRatingSummary struct {
	ID         uint    `json:"id"`
	Bitterness int     `json:"bitterness"`
	Acidity    int     `json:"acidity"`
	Body       int     `json:"body"`
	FlavorNote *string `json:"flavor_note,omitempty"`
}
