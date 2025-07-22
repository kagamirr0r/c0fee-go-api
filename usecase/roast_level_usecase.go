package usecase

import (
	"c0fee-api/model"
)

type RoastLevelUsecase interface {
	GetAllRoastLevels() []model.RoastLevelResponse
}

type roastLevelUsecase struct{}

func (rlu *roastLevelUsecase) GetAllRoastLevels() []model.RoastLevelResponse {
	return model.GetAllRoastLevels()
}

func NewRoastLevelUsecase() RoastLevelUsecase {
	return &roastLevelUsecase{}
}
