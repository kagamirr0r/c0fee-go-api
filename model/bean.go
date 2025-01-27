package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bean struct {
	gorm.Model
	ID         uint           `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	Name       *string        `json:"name"`
	UserID     uuid.UUID      `json:"user_id" gorm:"type:uuid;" validate:"required"`
	User       User           `json:"user"`
	RoasterID  uint           `json:"roaster_id" validate:"required"`
	Roaster    Roaster        `json:"roaster"`
	Countries  []Country      `json:"countries" gorm:"many2many:bean_countries;"`
	Varieties  []Variety      `json:"varieties" gorm:"many2many:bean_varieties;"`
	Area       *string        `json:"area" validate:"required"`
	RoastLevel string         `json:"roast_level" validate:"required"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type BeanResponse struct {
	Beans []Bean `json:"beans"`
}
