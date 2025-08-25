package model

import (
	"time"

	"gorm.io/gorm"
)

type Area struct {
	ID        uint   `param:"id" gorm:"primary_key;" validate:"required"`
	Name      string `gorm:"unique" validate:"required"`
	CountryID uint   `gorm:"not null" validate:"required"`
	Country   Country
	Farms     []Farm `gorm:"foreignKey:AreaID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
