package model

import (
	"time"

	"gorm.io/gorm"
)

type Variety struct {
	ID        uint           `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	Name      string         `json:"name" gorm:"unique" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type VarietyListResponse struct {
	ID   uint   `json:"id" param:"id"`
	Name string `json:"name"`
}

func (v *Variety) ToListResponse() VarietyListResponse {
	return VarietyListResponse{
		ID:   v.ID,
		Name: v.Name,
	}
}

type VarietiesResponse struct {
	Varieties []VarietyListResponse `json:"varieties"`
	Count     uint                  `json:"count"`
}
