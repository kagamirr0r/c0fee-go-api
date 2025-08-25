package usecase

import (
	"c0fee-api/dto"
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IRoastLevelUsecase interface {
	GetAll() ([]dto.IdNameSummary, error)
}

type roastLevelUsecase struct {
	rlr repository.IRoastLevelRepository
}

func (rlu *roastLevelUsecase) GetAll() ([]dto.IdNameSummary, error) {
	var roastLevels []model.RoastLevel
	if err := rlu.rlr.GetAll(&roastLevels); err != nil {
		return nil, err
	}

	var result []dto.IdNameSummary
	for _, rl := range roastLevels {
		result = append(result, dto.IdNameSummary{
			ID:   rl.ID,
			Name: rl.Name,
		})
	}

	return result, nil
}

func NewRoastLevelUsecase(rlr repository.IRoastLevelRepository) IRoastLevelUsecase {
	return &roastLevelUsecase{rlr}
}
