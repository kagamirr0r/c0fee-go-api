package usecase

import (
	"c0fee-api/domain/entity"
	domainRepo "c0fee-api/domain/repository"
	"c0fee-api/dto"
)

type IFarmUsecase interface {
	Read(id uint) (dto.FarmOutput, error)
}

type farmUsecase struct {
	ar domainRepo.IFarmRepository
}

func (fu *farmUsecase) Read(id uint) (dto.FarmOutput, error) {
	var storedFarm entity.Farm
	if err := fu.ar.GetById(&storedFarm, id); err != nil {
		return dto.FarmOutput{}, err
	}

	return fu.convertToFarmResponse(&storedFarm), nil
}

func (fu *farmUsecase) convertToFarmResponse(farm *entity.Farm) dto.FarmOutput {
	farmers := make([]dto.FarmerListOutput, len(farm.Farmers))
	for i, farmer := range farm.Farmers {
		farmers[i] = dto.FarmerListOutput{
			ID:   farmer.ID,
			Name: farmer.Name,
		}
	}

	return dto.FarmOutput{
		ID:      farm.ID,
		Name:    farm.Name,
		Farmers: farmers,
	}
}

func NewFarmUsecase(ar domainRepo.IFarmRepository) IFarmUsecase {
	return &farmUsecase{ar}
}
