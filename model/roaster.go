package model

import (
	"time"

	"gorm.io/gorm"
)

type Roaster struct {
	ID        uint   `param:"id" gorm:"primary_key;" validate:"required"`
	Name      string `gorm:"unique" validate:"required"`
	Address   string `validate:"required"`
	WebURL    string `validate:"required"`
	Beans     []Bean `gorm:"hasMany:Beans;foreignKey:RoasterID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
