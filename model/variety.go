package model

import (
	"time"

	"gorm.io/gorm"
)

type Variety struct {
	ID        uint   `param:"id" gorm:"primary_key;"`
	Name      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
