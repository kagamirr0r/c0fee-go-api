package usecase

import (
	"c0fee-api/model"
	"c0fee-api/repository"
	"context"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
)

type IBeanUsecase interface {
	Read(bean model.Bean) (model.BeanResponse, error)
	ListByUser(user model.User) (model.BeansResponse, error)
}

type beanUsecase struct {
	ur       repository.IUserRepository
	br       repository.IBeanRepository
	s3Client *minio.Client
}

func (bu *beanUsecase) generatePresignedURL(imageKey string) (string, error) {
	if imageKey == "" || imageKey == "null" {
		return "", nil
	}

	presignedURL, err := bu.s3Client.PresignedGetObject(
		context.Background(),
		os.Getenv("S3_BUCKET"),
		"beans/"+imageKey,
		time.Hour*1,
		nil,
	)

	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

func (bu *beanUsecase) Read(bean model.Bean) (model.BeanResponse, error) {
	storedBean := model.Bean{}
	if err := bu.br.GetBeanById(&storedBean, bean.ID); err != nil {
		return model.BeanResponse{}, err
	}

	imageURL, err := bu.generatePresignedURL(storedBean.ImageKey)
	if err != nil {
		return model.BeanResponse{}, err
	}

	return model.BeanResponse{
		ID:            storedBean.ID,
		Name:          storedBean.Name,
		User:          storedBean.User,
		Roaster:       storedBean.Roaster,
		ProcessMethod: storedBean.ProcessMethod,
		Countries:     storedBean.Countries,
		Varieties:     storedBean.Varieties,
		Area:          storedBean.Area,
		RoastLevel:    storedBean.RoastLevel,
		ImageURL:      imageURL,
	}, nil
}

func (bu *beanUsecase) ListByUser(user model.User) (model.BeansResponse, error) {
	targetUser := model.User{}
	if err := bu.ur.GetUserById(&targetUser, user.ID); err != nil {
		return model.BeansResponse{}, err
	}

	beans, err := bu.br.GetBeansByUserId(targetUser.ID)
	if err != nil {
		return model.BeansResponse{}, err
	}

	// Convert []Bean to []BeanResponse
	beanResponses := make([]model.BeanResponse, len(beans))
	for i, bean := range beans {
		imageURL, err := bu.generatePresignedURL(bean.ImageKey)
		if err != nil {
			return model.BeansResponse{}, err
		}
		beanResponses[i] = model.BeanResponse{
			ID:            bean.ID,
			Name:          bean.Name,
			User:          bean.User,
			Roaster:       bean.Roaster,
			ProcessMethod: bean.ProcessMethod,
			Countries:     bean.Countries,
			Varieties:     bean.Varieties,
			Area:          bean.Area,
			RoastLevel:    bean.RoastLevel,
			ImageURL:      imageURL,
		}
	}

	return model.BeansResponse{Beans: beanResponses}, nil
}

func NewBeanUsecase(ur repository.IUserRepository, br repository.IBeanRepository, s3Client *minio.Client) IBeanUsecase {
	return &beanUsecase{ur, br, s3Client}
}
