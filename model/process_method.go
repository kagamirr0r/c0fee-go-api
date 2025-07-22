package model

import (
	"time"

	"gorm.io/gorm"
)

type ProcessMethod struct {
	ID        uint           `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	Name      string         `json:"name" gorm:"unique" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type ProcessMethodResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ProcessMethodsResponse struct {
	ProcessMethods []ProcessMethodResponse `json:"process_methods"`
	Count          uint                    `json:"count"`
}

func (pm *ProcessMethod) ToResponse() ProcessMethodResponse {
	return ProcessMethodResponse{
		ID:   pm.ID,
		Name: pm.Name,
	}
}
