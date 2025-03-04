package model

import (
	"time"

	"gorm.io/gorm"
)

type Roaster struct {
	ID        uint           `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	Name      string         `json:"name" gorm:"unique" validate:"required"`
	Address   string         `json:"address" validate:"required"`
	WebURL    string         `json:"web_url" validate:"required"`
	Beans     []Bean         `json:"beans" gorm:"hasMany:Beans;foreignKey:RoasterID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// type RoasterResponse struct {
// 	ID   uuid.UUID `json:"id" gorm:"primary_key;type:uuid;"`
// 	Name string    `json:"name" gorm:"unique" validate:"required"`
// }
