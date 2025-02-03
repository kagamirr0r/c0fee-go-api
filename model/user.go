package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID   uuid.UUID `json:"id" param:"id" gorm:"primary_key;type:uuid;" validate:"required"`
	Name string    `json:"name" gorm:"unique" validate:"required"`
	gorm.Model
}

type UserResponse struct {
	ID   uuid.UUID `json:"id" gorm:"primary_key;type:uuid;"`
	Name string    `json:"name" gorm:"unique" validate:"required"`
}
