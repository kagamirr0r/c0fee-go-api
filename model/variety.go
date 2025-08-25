package model

import (
	"time"

	"gorm.io/gorm"
)

type Variety struct {
	ID        uint   `param:"id" gorm:"primary_key;" validate:"required"`
	Name      string `gorm:"unique" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
