package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoastLevelType string

const (
	Light       RoastLevelType = "Light"
	MediumLight RoastLevelType = "Medium-Light"
	Medium      RoastLevelType = "Medium"
	MediumDark  RoastLevelType = "Medium-Dark"
	Dark        RoastLevelType = "Dark"
)

type Bean struct {
	ID              uint           `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	Name            string         `json:"name"`
	UserID          uuid.UUID      `json:"user_id" gorm:"type:uuid;not null" validate:"required"`
	User            User           `json:"user"`
	RoasterID       uint           `json:"roaster_id" gorm:"not null" validate:"required"`
	Roaster         Roaster        `json:"roaster"`
	ProcessMethodID uint           `json:"process_method_id" gorm:"not null" validate:"required"`
	ProcessMethod   ProcessMethod  `json:"process_method"`
	Countries       []Country      `json:"countries" gorm:"many2many:bean_countries;"`
	Varieties       []Variety      `json:"varieties" gorm:"many2many:bean_varieties;"`
	Area            string         `json:"area"`
	RoastLevel      RoastLevelType `json:"roast_level" gorm:"not null;default:Medium" validate:"required"`
	gorm.Model
}

type BeanResponse struct {
	Beans []Bean `json:"beans"`
}
