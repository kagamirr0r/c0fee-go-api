package usecase

import (
	"c0fee-api/common"
	"c0fee-api/common/converter"
	"c0fee-api/dto"
	"c0fee-api/infrastructure/s3"
	"c0fee-api/model"
	"c0fee-api/repository"
)

type IUserUsecase interface {
	Create(user model.User) (dto.UserOutput, error)
	Read(user model.User) (dto.UserOutput, error)
	GetUserBeans(user model.User, params common.QueryParams) (dto.UserBeansOutput, error)
}

type userUsecase struct {
	ur        repository.IUserRepository
	br        repository.IBeanRepository
	s3Service s3.IS3Service
}

func (uu *userUsecase) Create(user model.User) (dto.UserOutput, error) {
	newUser := model.User{ID: user.ID, Name: user.Name}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return dto.UserOutput{}, err
	}
	return dto.UserOutput{ID: newUser.ID.String(), Name: newUser.Name}, nil
}

func (uu *userUsecase) Read(user model.User) (dto.UserOutput, error) {
	storedUser := model.User{}
	if err := uu.ur.GetById(&storedUser, user.ID); err != nil {
		return dto.UserOutput{}, err
	}

	var avatarURL string
	if storedUser.AvatarKey != "" {
		presignedURL, err := uu.s3Service.GenerateUserAvatarURL(storedUser.AvatarKey)
		if err != nil {
			return dto.UserOutput{}, err
		}
		avatarURL = presignedURL
	}
	return dto.UserOutput{ID: storedUser.ID.String(), Name: storedUser.Name, AvatarURL: avatarURL}, nil
}

func (uu *userUsecase) GetUserBeans(user model.User, params common.QueryParams) (dto.UserBeansOutput, error) {
	storedUser := model.User{}
	if err := uu.ur.GetById(&storedUser, user.ID); err != nil {
		return dto.UserBeansOutput{}, err
	}

	beans := []model.Bean{}

	// パラメータが存在する場合は検索を使用、そうでなければリスト全体を取得
	if params.NameLike != "" || params.Limit > 0 {
		err := uu.br.SearchBeansByUserId(&beans, storedUser.ID, params)
		if err != nil {
			return dto.UserBeansOutput{}, err
		}
	} else {
		err := uu.br.GetBeansByUserId(&beans, storedUser.ID)
		if err != nil {
			return dto.UserBeansOutput{}, err
		}
	}

	userResponse := dto.UserOutput{
		ID:   storedUser.ID.String(),
		Name: storedUser.Name,
	}

	beansResponse := make([]dto.BeanOutput, len(beans))
	for i, bean := range beans {
		var imageURL string
		if bean.ImageKey != nil {
			url, err := uu.s3Service.GenerateBeanImageURL(*bean.ImageKey)
			if err != nil {
				return dto.UserBeansOutput{}, err
			}
			imageURL = url
		}
		beansResponse[i] = converter.ConvertToBeanResponse(&bean, imageURL)
	}

	return dto.UserBeansOutput{
		User:  userResponse,
		Beans: beansResponse,
		Count: uint(len(beans)),
	}, nil
}

func NewUserUsecase(ur repository.IUserRepository, br repository.IBeanRepository, s3Service s3.IS3Service) IUserUsecase {
	return &userUsecase{ur, br, s3Service}
}
