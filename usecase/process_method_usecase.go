package usecase

import (
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IProcessMethodUsecase interface {
	List() (model.ProcessMethodsResponse, error)
}

type processMethodUsecase struct {
	pmr repository.IProcessMethodRepository
}

func (pmu *processMethodUsecase) List() (model.ProcessMethodsResponse, error) {
	processMethods := []model.ProcessMethod{}
	err := pmu.pmr.List(&processMethods)
	if err != nil {
		return model.ProcessMethodsResponse{}, err
	}

	processMethodResponses := make([]model.ProcessMethodResponse, len(processMethods))
	for i, processMethod := range processMethods {
		processMethodResponses[i] = processMethod.ToResponse()
	}

	return model.ProcessMethodsResponse{ProcessMethods: processMethodResponses, Count: uint(len(processMethods))}, nil
}

func NewProcessMethodUsecase(pmr repository.IProcessMethodRepository) IProcessMethodUsecase {
	return &processMethodUsecase{pmr}
}
