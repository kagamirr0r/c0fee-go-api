package usecase

import (
	"c0fee-api/domain/farm"
	"c0fee-api/dto"
)

type IFarmUsecase interface {
	Read(id uint) (dto.FarmOutput, error)
}

type farmUsecase struct {
	ar farm.IFarmRepository
}

func (fu *farmUsecase) Read(id uint) (dto.FarmOutput, error) {
	var storedFarm farm.Entity
	if err := fu.ar.GetById(&storedFarm, id); err != nil {
		return dto.FarmOutput{}, err
	}

	return fu.convertToFarmResponse(&storedFarm), nil
}

func (fu *farmUsecase) convertToFarmResponse(farmEntity *farm.Entity) dto.FarmOutput {
	farmers := make([]dto.FarmerListOutput, len(farmEntity.Farmers))
	for i, farmer := range farmEntity.Farmers {
		farmers[i] = dto.FarmerListOutput{
			ID:   farmer.ID,
			Name: farmer.Name,
		}
	}

	return dto.FarmOutput{
		ID:      farmEntity.ID,
		Name:    farmEntity.Name,
		Farmers: farmers,
	}
}

func NewFarmUsecase(ar farm.IFarmRepository) IFarmUsecase {
	return &farmUsecase{ar}
}
