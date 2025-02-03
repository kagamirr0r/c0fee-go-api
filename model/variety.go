package model

import (
	"gorm.io/gorm"
)

type Variety struct {
	ID      uint   `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	Variety string `json:"variety" validate:"required"`
	gorm.Model
}
