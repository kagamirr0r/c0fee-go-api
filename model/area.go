package model

import (
	"time"

	"gorm.io/gorm"
)

type Area struct {
	ID        uint           `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	Name      string         `json:"name" gorm:"unique" validate:"required"`
	CountryID uint           `json:"country_id" gorm:"not null" validate:"required"`
	Country   Country        `json:"country"`
	Farms     []Farm         `json:"farms" gorm:"foreignKey:AreaID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
