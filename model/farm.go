package model

import (
	"time"

	"gorm.io/gorm"
)

type Farm struct {
	ID        uint   `param:"id" gorm:"primary_key;"`
	Name      string `gorm:"not null"`
	AreaID    uint   `gorm:"not null"`
	Area      Area
	Farmers   []Farmer `gorm:"foreignKey:FarmID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
