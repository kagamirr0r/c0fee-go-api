package usecase

import (
	"c0fee-api/domain/entity"
	domainRepo "c0fee-api/domain/repository"
	"c0fee-api/dto"
)

type IProcessMethodUsecase interface {
	List() (dto.ProcessMethodsOutput, error)
}

type processMethodUsecase struct {
	pmr domainRepo.IProcessMethodRepository
}

func (pmu *processMethodUsecase) List() (dto.ProcessMethodsOutput, error) {
	var processMethods []entity.ProcessMethod
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

func NewProcessMethodUsecase(pmr domainRepo.IProcessMethodRepository) IProcessMethodUsecase {
	return &processMethodUsecase{pmr}
}
