package model

import (
	"time"

	"gorm.io/gorm"
)

type Country struct {
	ID        uint   `param:"id" gorm:"primary_key;" validate:"required"`
	Name      string `gorm:"unique" validate:"required"`
	Code      string `gorm:"unique" validate:"required"`
	Areas     []Area `gorm:"foreignKey:CountryID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
