package model

import (
	"time"

	"gorm.io/gorm"
)

type ProcessMethod struct {
	ID        uint           `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	Method    string         `json:"method" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
