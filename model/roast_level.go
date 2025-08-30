package model

import (
	"time"

	"gorm.io/gorm"
)

type RoastLevel struct {
	ID        uint   `gorm:"primary_key;"`
	Name      string `gorm:"not null;unique"`
	Level     int    `gorm:"not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
