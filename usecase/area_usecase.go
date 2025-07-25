package usecase

import (
	"c0fee-api/dto"
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IAreaUsecase interface {
	Read(id uint) (dto.AreaResponse, error)
}

type areaUsecase struct {
	ar repository.IAreaRepository
}

func (au *areaUsecase) Read(id uint) (dto.AreaResponse, error) {
	storedArea := model.Area{}
	if err := au.ar.GetById(&storedArea, id); err != nil {
		return dto.AreaResponse{}, err
	}

	return au.convertToAreaResponse(&storedArea), nil
}

func (au *areaUsecase) convertToAreaResponse(area *model.Area) dto.AreaResponse {
	farms := make([]dto.FarmListResponse, len(area.Farms))
	for i, farm := range area.Farms {
		farms[i] = dto.FarmListResponse{
			ID:   farm.ID,
			Name: farm.Name,
		}
	}

	return dto.AreaResponse{
		ID:    area.ID,
		Name:  area.Name,
		Farms: farms,
	}
}

func NewAreaUsecase(ar repository.IAreaRepository) IAreaUsecase {
	return &areaUsecase{ar}
}
