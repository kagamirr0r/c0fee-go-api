package usecase

import (
	"c0fee-api/dto"
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IVarietyUsecase interface {
	List() (dto.VarietiesOutput, error)
}

type varietyUsecase struct {
	vr repository.IVarietyRepository
}

func (vu *varietyUsecase) List() (dto.VarietiesOutput, error) {
	varieties := []model.Variety{}
	err := vu.vr.List(&varieties)
	if err != nil {
		return dto.VarietiesOutput{}, err
	}

	varietyResponses := make([]dto.VarietyListOutput, len(varieties))
	for i, variety := range varieties {
		varietyResponses[i] = dto.VarietyListOutput{
			ID:   variety.ID,
			Name: variety.Name,
		}
	}

	return dto.VarietiesOutput{Varieties: varietyResponses, Count: uint(len(varieties))}, nil
}

func NewVarietyUsecase(vr repository.IVarietyRepository) IVarietyUsecase {
	return &varietyUsecase{vr}
}
