package usecase

import (
	"c0fee-api/domain/process_method"
	"c0fee-api/dto"
)

type IProcessMethodUsecase interface {
	List() (dto.ProcessMethodsOutput, error)
}

type processMethodUsecase struct {
	pmr process_method.IProcessMethodRepository
}

func (pmu *processMethodUsecase) List() (dto.ProcessMethodsOutput, error) {
	var processMethods []process_method.Entity
	err := pmu.pmr.List(&processMethods)
	if err != nil {
		return dto.ProcessMethodsOutput{}, err
	}

	processMethodResponses := make([]dto.ProcessMethodOutput, len(processMethods))
	for i, processMethodEntity := range processMethods {
		processMethodResponses[i] = dto.ProcessMethodOutput{
			ID:   processMethodEntity.ID,
			Name: processMethodEntity.Name,
		}
	}

	return dto.ProcessMethodsOutput{ProcessMethods: processMethodResponses, Count: uint(len(processMethods))}, nil
}

func NewProcessMethodUsecase(pmr process_method.IProcessMethodRepository) IProcessMethodUsecase {
	return &processMethodUsecase{pmr}
}
