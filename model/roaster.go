package model

import (
	"gorm.io/gorm"
)

type Roaster struct {
	ID      uint   `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	WebURL  string `json:"web_url" validate:"required"`
	gorm.Model
}

// type RoasterResponse struct {
// 	ID   uuid.UUID `json:"id" gorm:"primary_key;type:uuid;"`
// 	Name string    `json:"name" gorm:"unique" validate:"required"`
// }
