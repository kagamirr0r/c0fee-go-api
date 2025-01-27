package usecase

import (
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IUserUsecase interface {
	Create(user model.User) (model.UserResponse, error)
	Show(user model.User) (model.UserResponse, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func (uu *userUsecase) Create(user model.User) (model.UserResponse, error) {
	newUser := model.User{ID: user.ID, Name: user.Name}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	return model.UserResponse{ID: newUser.ID, Name: newUser.Name}, nil
}

func (uu *userUsecase) Show(user model.User) (model.UserResponse, error) {
	storedUser := model.User{}
	if err := uu.ur.GetUserById(&storedUser, user.ID); err != nil {
		return model.UserResponse{}, err
	}
	return model.UserResponse{ID: storedUser.ID, Name: storedUser.Name}, nil
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}
