package usecase

import (
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IFarmUsecase interface {
	Read(id uint) (model.FarmResponse, error)
}

type farmUsecase struct {
	ar repository.IFarmRepository
}

func (au *farmUsecase) Read(id uint) (model.FarmResponse, error) {
	storedFarm := model.Farm{}
	if err := au.ar.GetById(&storedFarm, id); err != nil {
		return model.FarmResponse{}, err
	}

	return storedFarm.ToResponse(), nil
}

func NewFarmUsecase(ar repository.IFarmRepository) IFarmUsecase {
	return &farmUsecase{ar}
}
