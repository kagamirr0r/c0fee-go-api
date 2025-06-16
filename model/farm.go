package model

import (
	"time"

	"gorm.io/gorm"
)

type Farm struct {
	ID        uint           `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	Name      string         `json:"name" gorm:"not null" validate:"required"`
	AreaID    uint           `json:"area_id" gorm:"not null" validate:"required"`
	Area      Area           `json:"area"`
	Farmers   []Farmer       `json:"farmers" gorm:"foreignKey:FarmID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type FarmResponse struct {
	ID      uint                 `json:"id"`
	Name    string               `json:"name"`
	Farmers []FarmerListResponse `json:"farmers"`
}

func (f *Farm) ToResponse() FarmResponse {
	farmers := make([]FarmerListResponse, len(f.Farmers))
	for i, farmer := range f.Farmers {
		farmers[i] = farmer.ToListResponse()
	}

	return FarmResponse{
		ID:      f.ID,
		Name:    f.Name,
		Farmers: farmers,
	}
}

type FarmListResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type FarmsResponse struct {
	Areas []AreaListResponse `json:"areas"`
	Count uint               `json:"count"`
}

func (a *Farm) ToListResponse() FarmListResponse {
	return FarmListResponse{
		ID:   a.ID,
		Name: a.Name,
	}
}
