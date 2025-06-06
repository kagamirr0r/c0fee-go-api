package usecase

import (
	"c0fee-api/infrastructure/s3"
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IBeanUsecase interface {
	Read(bean model.Bean) (model.BeanResponse, error)
}

type beanUsecase struct {
	ur        repository.IUserRepository
	br        repository.IBeanRepository
	s3Service s3.IS3Service
}

func (bu *beanUsecase) Read(bean model.Bean) (model.BeanResponse, error) {
	storedBean := model.Bean{}
	if err := bu.br.GetBeanById(&storedBean, bean.ID); err != nil {
		return model.BeanResponse{}, err
	}

	imageURL, err := bu.s3Service.GenerateBeanImageURL(*storedBean.ImageKey)
	if err != nil {
		return model.BeanResponse{}, err
	}

	return storedBean.ToResponse(imageURL), nil
}

func NewBeanUsecase(ur repository.IUserRepository, br repository.IBeanRepository, s3Service s3.IS3Service) IBeanUsecase {
	return &beanUsecase{ur, br, s3Service}
}
