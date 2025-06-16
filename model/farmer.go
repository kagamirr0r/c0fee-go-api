package model

import (
	"time"

	"gorm.io/gorm"
)

type Farmer struct {
	ID        uint           `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	Name      string         `json:"name" gorm:"not null" validate:"required"`
	FarmID    uint           `json:"farm_id" gorm:"not null" validate:"required"`
	Farm      Farm           `json:"farm"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type FarmerListResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (f *Farmer) ToListResponse() FarmerListResponse {
	return FarmerListResponse{
		ID:   f.ID,
		Name: f.Name,
	}
}
