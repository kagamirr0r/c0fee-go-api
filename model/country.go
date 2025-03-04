package model

import (
	"time"

	"gorm.io/gorm"
)

type Country struct {
	ID        uint           `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	Name      string         `json:"name" gorm:"unique" validate:"required"`
	Code      string         `json:"code" gorm:"unique" validate:"required"`
	Areas     []Area         `json:"areas" gorm:"foreignKey:CountryID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
