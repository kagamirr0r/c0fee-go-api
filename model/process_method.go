package model

import (
	"gorm.io/gorm"
)

type ProcessMethod struct {
	ID     uint   `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	Method string `json:"method" validate:"required"`
	gorm.Model
}
