package usecase

import (
	"c0fee-api/infrastructure/s3"
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IUserUsecase interface {
	Create(user model.User) (model.UserResponse, error)
	Read(user model.User) (model.UserResponse, error)
}

type userUsecase struct {
	ur        repository.IUserRepository
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
	if err := uu.ur.GetUserById(&storedUser, user.ID); err != nil {
		return model.UserResponse{}, err
	}

	var avatarURL string
	if storedUser.AvatarKey != "" {
		// ユースケース層で presigned URL を生成
		presignedURL, err := uu.s3Service.GenerateUserAvatarURL(storedUser.AvatarKey)

		if err != nil {
			return model.UserResponse{}, err
		}
		avatarURL = presignedURL
	}
	return model.UserResponse{ID: storedUser.ID, Name: storedUser.Name, AvatarURL: avatarURL}, nil
}

func NewUserUsecase(ur repository.IUserRepository, s3Service s3.IS3Service) IUserUsecase {
	return &userUsecase{ur, s3Service}
}
