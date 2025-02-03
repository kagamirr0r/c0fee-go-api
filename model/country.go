package model

import (
	"gorm.io/gorm"
)

type Country struct {
	ID   uint   `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
	gorm.Model
}
