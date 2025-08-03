package usecase

import (
	"c0fee-api/dto"
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IProcessMethodUsecase interface {
	List() (dto.ProcessMethodsOutput, error)
}

type processMethodUsecase struct {
	pmr repository.IProcessMethodRepository
}

func (pmu *processMethodUsecase) List() (dto.ProcessMethodsOutput, error) {
	processMethods := []model.ProcessMethod{}
	err := pmu.pmr.List(&processMethods)
	if err != nil {
		return dto.ProcessMethodsOutput{}, err
	}

	processMethodResponses := make([]dto.ProcessMethodOutput, len(processMethods))
	for i, processMethod := range processMethods {
		processMethodResponses[i] = dto.ProcessMethodOutput{
			ID:   processMethod.ID,
			Name: processMethod.Name,
		}
	}

	return dto.ProcessMethodsOutput{ProcessMethods: processMethodResponses, Count: uint(len(processMethods))}, nil
}

func NewProcessMethodUsecase(pmr repository.IProcessMethodRepository) IProcessMethodUsecase {
	return &processMethodUsecase{pmr}
}
