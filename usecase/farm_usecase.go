package usecase

import (
	"c0fee-api/dto"
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IFarmUsecase interface {
	Read(id uint) (dto.FarmOutput, error)
}

type farmUsecase struct {
	ar repository.IFarmRepository
}

func (au *farmUsecase) Read(id uint) (dto.FarmOutput, error) {
	storedFarm := model.Farm{}
	if err := au.ar.GetById(&storedFarm, id); err != nil {
		return dto.FarmOutput{}, err
	}

	return au.convertToFarmResponse(&storedFarm), nil
}

func (au *farmUsecase) convertToFarmResponse(farm *model.Farm) dto.FarmOutput {
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

func NewFarmUsecase(ar repository.IFarmRepository) IFarmUsecase {
	return &farmUsecase{ar}
}
