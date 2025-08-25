package model

import (
	"time"

	"gorm.io/gorm"
)

type Farmer struct {
	ID        uint   `param:"id" gorm:"primary_key;" validate:"required"`
	Name      string `gorm:"not null" validate:"required"`
	FarmID    uint   `gorm:"not null" validate:"required"`
	Farm      Farm
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
