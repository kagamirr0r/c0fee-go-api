package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID      `json:"id" param:"id" gorm:"primary_key;type:uuid;" validate:"required"`
	Name      string         `json:"name" gorm:"unique" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserResponse struct {
	ID   uuid.UUID `json:"id" gorm:"primary_key;type:uuid;"`
	Name string    `json:"name" gorm:"unique" validate:"required"`
}
