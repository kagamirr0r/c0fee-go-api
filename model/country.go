package model

import (
	"time"

	"gorm.io/gorm"
)

type Country struct {
	ID        uint           `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	Name      string         `json:"name" gorm:"unique" validate:"required"`
	Code      string         `json:"code" gorm:"unique" validate:"required"`
	Areas     []Area         `json:"areas" gorm:"foreignKey:CountryID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// 詳細情報を返すためのレスポンス構造体
type CountryResponse struct {
	ID    uint               `json:"id"`
	Name  string             `json:"name"`
	Code  string             `json:"code"`
	Areas []AreaListResponse `json:"areas"`
}

func (c *Country) ToResponse() CountryResponse {
	areas := make([]AreaListResponse, len(c.Areas))
	for i, area := range c.Areas {
		areas[i] = area.ToListResponse()
	}

	return CountryResponse{
		ID:    c.ID,
		Name:  c.Name,
		Code:  c.Code,
		Areas: areas,
	}
}

// リストで返すためのレスポンス構造体
type CountryListResponse struct {
	ID   uint   `json:"id" param:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func (c *Country) ToListResponse() CountryListResponse {
	return CountryListResponse{
		ID:   c.ID,
		Name: c.Name,
		Code: c.Code,
	}
}

type CountriesResponse struct {
	Countries []CountryListResponse `json:"countries"`
	Count     uint                  `json:"count"`
}
