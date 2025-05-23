package usecase

import (
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IRoasterUsecase interface {
	List() (model.RoastersResponse, error)
}

type roasterUsecase struct {
	cr repository.IRoasterRepository
}

func (cu *roasterUsecase) List() (model.RoastersResponse, error) {
	roasters := []model.Roaster{}
	err := cu.cr.List(&roasters)
	if err != nil {
		return model.RoastersResponse{}, err
	}

	roastersResponse := make([]model.RoasterResponse, len(roasters))
	for i, roaster := range roasters {
		roastersResponse[i] = model.RoasterResponse{
			ID:      roaster.ID,
			Name:    roaster.Name,
			Address: roaster.Address,
			WebURL:  roaster.WebURL,
		}
	}

	return model.RoastersResponse{Roasters: roastersResponse, Count: uint(len(roasters))}, nil
}

func NewRoasterUsecase(cr repository.IRoasterRepository) IRoasterUsecase {
	return &roasterUsecase{cr}
}
