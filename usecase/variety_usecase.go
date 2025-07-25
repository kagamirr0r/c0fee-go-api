package usecase

import (
	"c0fee-api/dto"
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IVarietyUsecase interface {
	List() (dto.VarietiesResponse, error)
}

type varietyUsecase struct {
	vr repository.IVarietyRepository
}

func (vu *varietyUsecase) List() (dto.VarietiesResponse, error) {
	varieties := []model.Variety{}
	err := vu.vr.List(&varieties)
	if err != nil {
		return dto.VarietiesResponse{}, err
	}

	varietyResponses := make([]dto.VarietyListResponse, len(varieties))
	for i, variety := range varieties {
		varietyResponses[i] = dto.VarietyListResponse{
			ID:   variety.ID,
			Name: variety.Name,
		}
	}

	return dto.VarietiesResponse{Varieties: varietyResponses, Count: uint(len(varieties))}, nil
}

func NewVarietyUsecase(vr repository.IVarietyRepository) IVarietyUsecase {
	return &varietyUsecase{vr}
}
