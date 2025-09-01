package usecase

import (
	"c0fee-api/domain/entity"
	domainRepo "c0fee-api/domain/repository"
	"c0fee-api/dto"
)

type IAreaUsecase interface {
	Read(id uint) (dto.AreaOutput, error)
}

type areaUsecase struct {
	ar domainRepo.IAreaRepository
}

func (au *areaUsecase) Read(id uint) (dto.AreaOutput, error) {
	var storedArea entity.Area
	if err := au.ar.GetById(&storedArea, id); err != nil {
		return dto.AreaOutput{}, err
	}

	return au.convertToAreaResponse(&storedArea), nil
}

func (au *areaUsecase) convertToAreaResponse(area *entity.Area) dto.AreaOutput {
	farms := make([]dto.FarmListOutput, len(area.Farms))
	for i, farm := range area.Farms {
		farms[i] = dto.FarmListOutput{
			ID:   farm.ID,
			Name: farm.Name,
		}
	}

	return dto.AreaOutput{
		ID:    area.ID,
		Name:  area.Name,
		Farms: farms,
	}
}

func NewAreaUsecase(ar domainRepo.IAreaRepository) IAreaUsecase {
	return &areaUsecase{ar}
}
