package usecase

import (
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IBeanUsecase interface {
	ListByUser(user model.User) (model.BeansResponse, error)
}

type beanUsecase struct {
	ur repository.IUserRepository
	br repository.IBeanRepository
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
		}
	}

	return model.BeansResponse{Beans: beanResponses}, nil
}

func NewBeanUsecase(ur repository.IUserRepository, br repository.IBeanRepository) IBeanUsecase {
	return &beanUsecase{ur, br}
}
