package usecase

import (
	"c0fee-api/dto"
	"c0fee-api/model"
)

type RoastLevelUsecase interface {
	GetAllRoastLevels() []dto.RoastLevelOutput
}

type roastLevelUsecase struct{}

func (rlu *roastLevelUsecase) GetAllRoastLevels() []dto.RoastLevelOutput {
	roastLevels := make([]dto.RoastLevelOutput, len(model.AllRoastLevels))

	for i, level := range model.AllRoastLevels {
		roastLevels[i] = dto.RoastLevelOutput{
			ID:   i + 1,
			Name: string(level),
		}
	}
	return roastLevels
}

func NewRoastLevelUsecase() RoastLevelUsecase {
	return &roastLevelUsecase{}
}
