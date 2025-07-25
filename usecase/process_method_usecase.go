package usecase

import (
	"c0fee-api/dto"
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IProcessMethodUsecase interface {
	List() (dto.ProcessMethodsResponse, error)
}

type processMethodUsecase struct {
	pmr repository.IProcessMethodRepository
}

func (pmu *processMethodUsecase) List() (dto.ProcessMethodsResponse, error) {
	processMethods := []model.ProcessMethod{}
	err := pmu.pmr.List(&processMethods)
	if err != nil {
		return dto.ProcessMethodsResponse{}, err
	}

	processMethodResponses := make([]dto.ProcessMethodResponse, len(processMethods))
	for i, processMethod := range processMethods {
		processMethodResponses[i] = dto.ProcessMethodResponse{
			ID:   processMethod.ID,
			Name: processMethod.Name,
		}
	}

	return dto.ProcessMethodsResponse{ProcessMethods: processMethodResponses, Count: uint(len(processMethods))}, nil
}

func NewProcessMethodUsecase(pmr repository.IProcessMethodRepository) IProcessMethodUsecase {
	return &processMethodUsecase{pmr}
}
