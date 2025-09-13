package usecase

import (
	"c0fee-api/domain/variety"
	"c0fee-api/dto"
)

type IVarietyUsecase interface {
	List() (dto.VarietiesOutput, error)
}

type varietyUsecase struct {
	vr variety.IVarietyRepository
}

func (vu *varietyUsecase) List() (dto.VarietiesOutput, error) {
	var varieties []variety.Entity
	err := vu.vr.List(&varieties)
	if err != nil {
		return dto.VarietiesOutput{}, err
	}

	varietyResponses := make([]dto.VarietyListOutput, len(varieties))
	for i, varietyEntity := range varieties {
		varietyResponses[i] = dto.VarietyListOutput{
			ID:   varietyEntity.ID,
			Name: varietyEntity.Name,
		}
	}

	return dto.VarietiesOutput{Varieties: varietyResponses, Count: uint(len(varieties))}, nil
}

func NewVarietyUsecase(vr variety.IVarietyRepository) IVarietyUsecase {
	return &varietyUsecase{vr}
}
