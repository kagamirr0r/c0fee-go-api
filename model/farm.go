package model

import (
	"time"

	"gorm.io/gorm"
)

type Farm struct {
	ID        uint   `param:"id" gorm:"primary_key;" validate:"required"`
	Name      string `gorm:"not null" validate:"required"`
	AreaID    uint   `gorm:"not null" validate:"required"`
	Area      Area
	Farmers   []Farmer `gorm:"foreignKey:FarmID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
