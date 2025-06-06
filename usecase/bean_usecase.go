package usecase

import (
	"c0fee-api/common"
	"c0fee-api/infrastructure/s3"
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IBeanUsecase interface {
	Read(bean model.Bean) (model.BeanResponse, error)
	ListByUser(user model.User) (model.BeansResponse, error)
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

	return model.BeanResponse{
		ID:            storedBean.ID,
		Name:          storedBean.Name,
		User:          storedBean.User,
		Roaster:       storedBean.Roaster,
		ProcessMethod: storedBean.ProcessMethod,
		Country:       storedBean.Country,
		Varieties:     storedBean.Varieties,
		RoastLevel:    storedBean.RoastLevel,
		Area:          storedBean.Area,
		Farm:          storedBean.Farm,
		Farmer:        storedBean.Farmer,
		BeanRatings:   storedBean.BeanRatings,
		ImageURL:      common.StoPoint(imageURL),
	}, nil
}

func (bu *beanUsecase) ListByUser(user model.User) (model.BeansResponse, error) {
	targetUser := model.User{}
	if err := bu.ur.GetUserById(&targetUser, user.ID); err != nil {
		return model.BeansResponse{}, err
	}

	beans := []model.Bean{}
	err := bu.br.GetBeansByUserId(beans, targetUser.ID)
	if err != nil {
		return model.BeansResponse{}, err
	}

	beanResponses := make([]model.BeanResponse, len(beans))
	for i, bean := range beans {
		imageURL, err := bu.s3Service.GenerateBeanImageURL(*bean.ImageKey)
		if err != nil {
			return model.BeansResponse{}, err
		}

		beanResponses[i] = model.BeanResponse{
			ID:            bean.ID,
			Name:          bean.Name,
			Roaster:       bean.Roaster,
			ProcessMethod: bean.ProcessMethod,
			Country:       bean.Country,
			Varieties:     bean.Varieties,
			Area:          bean.Area,
			Farm:          bean.Farm,
			Farmer:        bean.Farmer,
			RoastLevel:    bean.RoastLevel,
			ImageURL:      common.StoPoint(imageURL),
		}
	}

	return model.BeansResponse{Beans: beanResponses, Count: uint(len(beans))}, nil
}

func NewBeanUsecase(ur repository.IUserRepository, br repository.IBeanRepository, s3Service s3.IS3Service) IBeanUsecase {
	return &beanUsecase{ur, br, s3Service}
}
