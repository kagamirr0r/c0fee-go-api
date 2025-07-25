package usecase

import (
	"c0fee-api/common"
	"c0fee-api/dto"
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IRoasterUsecase interface {
	List(params common.QueryParams) (dto.RoastersResponse, error)
}

type roasterUsecase struct {
	rr repository.IRoasterRepository
}

func (ru *roasterUsecase) List(params common.QueryParams) (dto.RoastersResponse, error) {
	roasters := []model.Roaster{}

	// パラメータが存在する場合は検索を使用、そうでなければリスト全体を取得
	if params.NameLike != "" || params.Limit > 0 {
		err := ru.rr.Search(&roasters, params)
		if err != nil {
			return dto.RoastersResponse{}, err
		}
	} else {
		err := ru.rr.List(&roasters)
		if err != nil {
			return dto.RoastersResponse{}, err
		}
	}

	roastersResponse := make([]dto.RoasterResponse, len(roasters))
	for i, roaster := range roasters {
		roastersResponse[i] = dto.RoasterResponse{
			ID:      roaster.ID,
			Name:    roaster.Name,
			Address: roaster.Address,
			WebURL:  roaster.WebURL,
		}
	}

	return dto.RoastersResponse{Roasters: roastersResponse, Count: uint(len(roasters))}, nil
}

func NewRoasterUsecase(cr repository.IRoasterRepository) IRoasterUsecase {
	return &roasterUsecase{cr}
}
