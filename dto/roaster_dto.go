package dto

import (
	"mime/multipart"
)

// Input DTOs
type RoasterFormInput struct {
	Data      string                `form:"data" validate:"required"`
	ImageFile *multipart.FileHeader `form:"image"`
}

type RoasterInput struct {
	Name    string  `json:"name" validate:"required"`
	Address string  `json:"address" validate:"required"`
	WebURL  *string `json:"web_url,omitempty"`
}

// Output DTOs
type RoasterOutput struct {
	ID       uint        `json:"id"`
	Name     string      `json:"name"`
	Address  string      `json:"address"`
	WebURL   string      `json:"web_url"`
	ImageURL *string     `json:"image_url,omitempty"`
	Beans    BeansOutput `json:"beans,omitempty"`
}

type RoastersOutput struct {
	Roasters []RoasterOutput `json:"roasters"`
	Count    uint            `json:"count"`
}
