package model

import (
	"time"

	"gorm.io/gorm"
)

type Roaster struct {
	ID        uint   `param:"id" gorm:"primary_key;"`
	Name      string `gorm:"unique"`
	Address   string
	WebURL    string
	Beans     []Bean  `gorm:"hasMany:Beans;foreignKey:RoasterID"`
	ImageKey  *string `gorm:"default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
