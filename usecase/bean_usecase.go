package usecase

import (
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IBeanUsecase interface {
	ListByUser(user model.User) (model.BeanResponse, error)
}

type beanUsecase struct {
	ur repository.IUserRepository
	br repository.IBeanRepository
}

func (bu *beanUsecase) ListByUser(user model.User) (model.BeanResponse, error) {
	targetUser := model.User{}
	if err := bu.ur.GetUserById(&targetUser, user.ID); err != nil {
		return model.BeanResponse{}, err
	}

	beans, err := bu.br.GetBeansByUserId(targetUser.ID)
	if err != nil {
		return model.BeanResponse{}, err
	}
	return model.BeanResponse{Beans: beans}, nil
}

func NewBeanUsecase(ur repository.IUserRepository, br repository.IBeanRepository) IBeanUsecase {
	return &beanUsecase{ur, br}
}
