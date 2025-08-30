package model

import (
	"time"

	"gorm.io/gorm"
)

type Farmer struct {
	ID        uint   `param:"id" gorm:"primary_key;"`
	Name      string `gorm:"not null"`
	FarmID    uint   `gorm:"not null"`
	Farm      Farm
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
