package usecase

import (
	"c0fee-api/common"
	"c0fee-api/common/converter/dto_entity"
	"c0fee-api/domain/entity"
	domainRepo "c0fee-api/domain/repository"
	"c0fee-api/dto"
	"c0fee-api/infrastructure/s3"
)

type IRoasterUsecase interface {
	List(params common.QueryParams) (dto.RoastersOutput, error)
	GetById(id uint) (dto.RoasterOutput, error)
}

type roasterUsecase struct {
	rr        domainRepo.IRoasterRepository
	br        domainRepo.IBeanRepository
	s3Service s3.IS3Service
}

func (ru *roasterUsecase) List(params common.QueryParams) (dto.RoastersOutput, error) {
	var roasters []entity.Roaster

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
	for i, roaster := range roasters {
		var imageURL *string
		if roaster.ImageKey != nil && *roaster.ImageKey != "" {
			url, err := ru.s3Service.GenerateRoasterImageURL(*roaster.ImageKey)
			if err == nil && url != "" {
				imageURL = &url
			}
		}

		roastersResponse[i] = dto.RoasterOutput{
			ID:       roaster.ID,
			Name:     roaster.Name,
			Address:  roaster.Address,
			WebURL:   roaster.WebURL,
			ImageURL: imageURL,
		}
	}

	return dto.RoastersOutput{Roasters: roastersResponse, Count: uint(len(roasters))}, nil
}

func (ru *roasterUsecase) GetById(id uint) (dto.RoasterOutput, error) {
	var roaster entity.Roaster
	err := ru.rr.GetById(&roaster, id)
	if err != nil {
		return dto.RoasterOutput{}, err
	}

	var imageURL *string
	if roaster.ImageKey != nil && *roaster.ImageKey != "" {
		url, err := ru.s3Service.GenerateRoasterImageURL(*roaster.ImageKey)
		if err == nil && url != "" {
			imageURL = &url
		}
	}

	// Use the preloaded beans from roaster entity
	beansOutput, err := dto_entity.BeanEntitiesToBeansOutput(roaster.Beans, common.QueryParams{}, ru.s3Service)
	if err != nil {
		return dto.RoasterOutput{}, err
	}

	return dto.RoasterOutput{
		ID:       roaster.ID,
		Name:     roaster.Name,
		Address:  roaster.Address,
		WebURL:   roaster.WebURL,
		ImageURL: imageURL,
		Beans:    beansOutput,
	}, nil
}

func NewRoasterUsecase(rr domainRepo.IRoasterRepository, br domainRepo.IBeanRepository, s3Service s3.IS3Service) IRoasterUsecase {
	return &roasterUsecase{rr, br, s3Service}
}
