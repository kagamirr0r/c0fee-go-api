package usecase

import (
	"c0fee-api/domain/entity"
	domainRepo "c0fee-api/domain/repository"
	"c0fee-api/dto"
)

type IVarietyUsecase interface {
	List() (dto.VarietiesOutput, error)
}

type varietyUsecase struct {
	vr domainRepo.IVarietyRepository
}

func (vu *varietyUsecase) List() (dto.VarietiesOutput, error) {
	var varieties []entity.Variety
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

func NewVarietyUsecase(vr domainRepo.IVarietyRepository) IVarietyUsecase {
	return &varietyUsecase{vr}
}
