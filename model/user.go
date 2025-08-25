package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `param:"id" gorm:"primary_key;type:uuid;" validate:"required"`
	Name      string    `gorm:"unique" validate:"required"`
	AvatarKey string    `gorm:"default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
