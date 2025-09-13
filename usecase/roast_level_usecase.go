package usecase

import (
	"c0fee-api/domain/roast_level"
	"c0fee-api/dto"
)

type IRoastLevelUsecase interface {
	GetAll() ([]dto.IdNameSummary, error)
}

type roastLevelUsecase struct {
	rlr roast_level.IRoastLevelRepository
}

func (rlu *roastLevelUsecase) GetAll() ([]dto.IdNameSummary, error) {
	var roastLevels []roast_level.Entity
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

func NewRoastLevelUsecase(rlr roast_level.IRoastLevelRepository) IRoastLevelUsecase {
	return &roastLevelUsecase{rlr}
}
