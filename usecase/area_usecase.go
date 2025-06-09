package usecase

import (
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IAreaUsecase interface {
	Read(id uint) (model.AreaResponse, error)
}

type areaUsecase struct {
	ar repository.IAreaRepository
}

func (au *areaUsecase) Read(id uint) (model.AreaResponse, error) {
	storedArea := model.Area{}
	if err := au.ar.GetById(&storedArea, id); err != nil {
		return model.AreaResponse{}, err
	}

	return storedArea.ToResponse(), nil
}

func NewAreaUsecase(ar repository.IAreaRepository) IAreaUsecase {
	return &areaUsecase{ar}
}
