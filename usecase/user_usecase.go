package usecase

import (
	"c0fee-api/model"
	"c0fee-api/repository"
	"context"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
)

type IUserUsecase interface {
	Create(user model.User) (model.UserResponse, error)
	Read(user model.User) (model.UserResponse, error)
}

type userUsecase struct {
	ur       repository.IUserRepository
	s3Client *minio.Client
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
		presignedURL, err := uu.s3Client.PresignedGetObject(context.Background(),
			os.Getenv("S3_BUCKET"),
			"users/"+storedUser.AvatarKey,
			time.Hour*1,
			nil)

		if err != nil {
			return model.UserResponse{}, err
		}
		avatarURL = presignedURL.String()
	}
	return model.UserResponse{ID: storedUser.ID, Name: storedUser.Name, AvatarURL: avatarURL}, nil
}

func NewUserUsecase(ur repository.IUserRepository, s3Client *minio.Client) IUserUsecase {
	return &userUsecase{ur, s3Client}
}
