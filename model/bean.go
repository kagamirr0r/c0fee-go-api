package model

import (
	"time"

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
	ProcessMethodID uint           `json:"process_method_id" validate:"required"`
	ProcessMethod   ProcessMethod  `json:"process_method"`
	Countries       []Country      `json:"countries" gorm:"many2many:bean_countries;"`
	Varieties       []Variety      `json:"varieties" gorm:"many2many:bean_varieties;"`
	Area            string         `json:"area"`
	RoastLevel      RoastLevelType `json:"roast_level" gorm:"not null;default:Medium" validate:"required"`
	ImageKey        string         `json:"image_key" gorm:"default:null"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type BeanResponse struct {
	ID            uint           `json:"id" param:"id"`
	Name          string         `json:"name"`
	User          User           `json:"user"`
	Roaster       Roaster        `json:"roaster"`
	ProcessMethod ProcessMethod  `json:"process_method"`
	Countries     []Country      `json:"countries"`
	Varieties     []Variety      `json:"varieties"`
	Area          string         `json:"area"`
	RoastLevel    RoastLevelType `json:"roast_level"`
	ImageURL      string         `json:"image_url,omitempty"`
}

type BeansResponse struct {
	Beans []BeanResponse `json:"beans"`
}
