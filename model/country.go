package model

import (
	"time"

	"gorm.io/gorm"
)

type Country struct {
	ID        uint   `param:"id" gorm:"primary_key;"`
	Name      string `gorm:"unique"`
	Code      string `gorm:"unique"`
	Areas     []Area `gorm:"foreignKey:CountryID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
