package usecase

import (
	"c0fee-api/common"
	"c0fee-api/infrastructure/s3"
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IUserUsecase interface {
	Create(user model.User) (model.UserResponse, error)
	Read(user model.User) (model.UserResponse, error)
	GetUserBeans(user model.User, params common.QueryParams) (model.UserBeansResponse, error)
}

type userUsecase struct {
	ur        repository.IUserRepository
	br        repository.IBeanRepository
	s3Service s3.IS3Service
}

func (uu *userUsecase) Create(user model.User) (model.UserResponse, error) {
	newUser := model.User{ID: user.ID, Name: user.Name}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	return model.UserResponse{ID: newUser.ID, Name: newUser.Name}, nil
}

func (uu *userUsecase) Read(user model.User) (model.UserResponse, error) {
	storedUser := model.User{}
	if err := uu.ur.GetById(&storedUser, user.ID); err != nil {
		return model.UserResponse{}, err
	}

	var avatarURL string
	if storedUser.AvatarKey != "" {
		presignedURL, err := uu.s3Service.GenerateUserAvatarURL(storedUser.AvatarKey)

		if err != nil {
			return model.UserResponse{}, err
		}
		avatarURL = presignedURL
	}
	return model.UserResponse{ID: storedUser.ID, Name: storedUser.Name, AvatarURL: avatarURL}, nil
}

func (uu *userUsecase) GetUserBeans(user model.User, params common.QueryParams) (model.UserBeansResponse, error) {
	storedUser := model.User{}
	if err := uu.ur.GetById(&storedUser, user.ID); err != nil {
		return model.UserBeansResponse{}, err
	}

	beans := []model.Bean{}
	err := uu.br.GetBeansByUserId(&beans, storedUser.ID)
	if err != nil {
		return model.UserBeansResponse{}, err
	}

	var avatarURL string
	if storedUser.AvatarKey != "" {
		presignedURL, err := uu.s3Service.GenerateUserAvatarURL(storedUser.AvatarKey)
		if err != nil {
			return model.UserBeansResponse{}, err
		}
		avatarURL = presignedURL
	}

	userResponse := storedUser.ToResponse(avatarURL)
	beansResponse := make([]model.BeanResponse, len(beans))
	for i, bean := range beans {
		imageURL, err := uu.s3Service.GenerateBeanImageURL(*bean.ImageKey)
		if err != nil {
			return model.UserBeansResponse{}, err
		}
		beansResponse[i] = bean.ToResponse(imageURL)
	}

	return model.UserBeansResponse{
		User:  userResponse,
		Beans: beansResponse,
		Count: uint(len(beans)),
	}, nil
}

func NewUserUsecase(ur repository.IUserRepository, bu repository.IBeanRepository, s3Service s3.IS3Service) IUserUsecase {
	return &userUsecase{ur, bu, s3Service}
}
