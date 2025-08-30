package model

import (
	"time"

	"gorm.io/gorm"
)

type Area struct {
	ID        uint   `param:"id" gorm:"primary_key;"`
	Name      string `gorm:"unique"`
	CountryID uint   `gorm:"not null"`
	Country   Country
	Farms     []Farm `gorm:"foreignKey:AreaID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
