package usecase

import (
	"c0fee-api/dto"
	"c0fee-api/model"
)

type RoastLevelUsecase interface {
	GetAllRoastLevels() []dto.RoastLevelResponse
}

type roastLevelUsecase struct{}

func (rlu *roastLevelUsecase) GetAllRoastLevels() []dto.RoastLevelResponse {
	roastLevels := make([]dto.RoastLevelResponse, len(model.AllRoastLevels))

	for i, level := range model.AllRoastLevels {
		roastLevels[i] = dto.RoastLevelResponse{
			ID:   i + 1,
			Name: string(level),
		}
	}
	return roastLevels
}

func NewRoastLevelUsecase() RoastLevelUsecase {
	return &roastLevelUsecase{}
}
