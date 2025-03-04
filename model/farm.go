package model

import (
	"time"

	"gorm.io/gorm"
)

type Farm struct {
	ID        uint           `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	Name      string         `json:"name" gorm:"not null" validate:"required"`
	AreaID    uint           `json:"area_id" gorm:"not null" validate:"required"`
	Area      Area           `json:"area"`
	Farmers   []Farmer       `json:"farmers" gorm:"foreignKey:FarmID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
