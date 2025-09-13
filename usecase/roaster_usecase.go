package usecase

import (
	"c0fee-api/common"
	"c0fee-api/domain/roaster"
	"c0fee-api/dto"
	"c0fee-api/infrastructure/s3"
)

type IRoasterUsecase interface {
	List(params common.QueryParams) (dto.RoastersOutput, error)
}

type roasterUsecase struct {
	rr        roaster.IRoasterRepository
	s3Service s3.IS3Service
}

func (ru *roasterUsecase) List(params common.QueryParams) (dto.RoastersOutput, error) {
	var roasters []roaster.Entity

	// パラメータが存在する場合は検索を使用、そうでなければリスト全体を取得
	if params.NameLike != "" || params.Limit > 0 {
		err := ru.rr.Search(&roasters, params)
		if err != nil {
			return dto.RoastersOutput{}, err
		}
	} else {
		err := ru.rr.List(&roasters)
		if err != nil {
			return dto.RoastersOutput{}, err
		}
	}

	roastersResponse := make([]dto.RoasterOutput, len(roasters))
	for i, roasterEntity := range roasters {
		var imageURL *string
		if roasterEntity.ImageKey != nil && *roasterEntity.ImageKey != "" {
			url, err := ru.s3Service.GenerateRoasterImageURL(*roasterEntity.ImageKey)
			if err == nil && url != "" {
				imageURL = &url
			}
		}

		roastersResponse[i] = dto.RoasterOutput{
			ID:       roasterEntity.ID,
			Name:     roasterEntity.Name,
			Address:  roasterEntity.Address,
			WebURL:   roasterEntity.WebURL,
			ImageURL: imageURL,
		}
	}

	return dto.RoastersOutput{Roasters: roastersResponse, Count: uint(len(roasters))}, nil
}

func NewRoasterUsecase(rr roaster.IRoasterRepository, s3Service s3.IS3Service) IRoasterUsecase {
	return &roasterUsecase{rr, s3Service}
}
