package model

import (
	"time"

	"gorm.io/gorm"
)

type Area struct {
	ID        uint           `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	Name      string         `json:"name" gorm:"unique" validate:"required"`
	CountryID uint           `json:"country_id" gorm:"not null" validate:"required"`
	Country   Country        `json:"country"`
	Farms     []Farm         `json:"farms" gorm:"foreignKey:AreaID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
type AreaResponse struct {
	ID    uint               `json:"id"`
	Name  string             `json:"name"`
	Code  string             `json:"code"`
	Farms []FarmListResponse `json:"farms"`
}

func (a *Area) ToResponse() AreaResponse {
	farms := make([]FarmListResponse, len(a.Farms))
	for i, farm := range a.Farms {
		farms[i] = farm.ToListResponse()
	}

	return AreaResponse{
		ID:    a.ID,
		Name:  a.Name,
		Farms: farms,
	}
}

type AreaListResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (a *Area) ToListResponse() AreaListResponse {
	return AreaListResponse{
		ID:   a.ID,
		Name: a.Name,
	}
}

type AreasResponse struct {
	Areas []AreaListResponse `json:"areas"`
	Count uint               `json:"count"`
}
