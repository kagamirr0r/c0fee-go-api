package model

import (
	"time"

	"gorm.io/gorm"
)

type RoastLevel struct {
	ID        uint   `gorm:"primary_key;" validate:"required"`
	Name      string `gorm:"not null;unique" validate:"required"`
	Level     int    `gorm:"not null;unique" validate:"required,min=1,max=10"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
