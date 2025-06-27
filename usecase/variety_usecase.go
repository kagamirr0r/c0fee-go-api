package usecase

import (
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IVarietyUsecase interface {
	List() (model.VarietiesResponse, error)
}

type varietyUsecase struct {
	vr repository.IVarietyRepository
}

func (vu *varietyUsecase) List() (model.VarietiesResponse, error) {
	varieties := []model.Variety{}
	err := vu.vr.List(&varieties)
	if err != nil {
		return model.VarietiesResponse{}, err
	}

	varietyResponses := make([]model.VarietyListResponse, len(varieties))
	for i, variety := range varieties {
		varietyResponses[i] = variety.ToListResponse()
	}

	return model.VarietiesResponse{Varieties: varietyResponses, Count: uint(len(varieties))}, nil
}

func NewVarietyUsecase(vr repository.IVarietyRepository) IVarietyUsecase {
	return &varietyUsecase{vr}
}