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
	GetUserBeans(user model.User, params common.QueryParams) (dto.BeansOutput, error)
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

func (uu *userUsecase) GetUserBeans(user model.User, params common.QueryParams) (dto.BeansOutput, error) {
	storedUser := model.User{}
	if err := uu.ur.GetById(&storedUser, user.ID); err != nil {
		return dto.BeansOutput{}, err
	}

	beans := []model.Bean{}
	if params.NameLike != "" || params.Cursor > 0 {
		err := uu.br.SearchBeansByUserId(&beans, storedUser.ID, params)
		if err != nil {
			return dto.BeansOutput{}, err
		}
	} else {
		err := uu.br.GetBeansByUserId(&beans, storedUser.ID, params)
		if err != nil {
			return dto.BeansOutput{}, err
		}
	}

	userBeans := make([]dto.BeanSummary, len(beans))
	for i, bean := range beans {
		var imageURL string
		if bean.ImageKey != nil {
			url, err := uu.s3Service.GenerateBeanImageURL(*bean.ImageKey)
			if err != nil {
				return dto.BeansOutput{}, err
			}
			imageURL = url
		}
		userBeans[i] = converter.ConvertBeanToBeanSummary(&bean, imageURL)
	}

	// カーソルページネーション用の情報を生成
	var nextCursor *uint
	if len(beans) > 0 && params.Limit > 0 && len(beans) == params.Limit {
		// 最後のBeanのIDをnext_cursorとして設定
		lastBeanID := beans[len(beans)-1].ID
		nextCursor = &lastBeanID
	}

	beansResponse := dto.BeansOutput{
		Beans:      userBeans,
		Count:      uint(len(beans)),
		NextCursor: nextCursor,
	}

	return beansResponse, nil
}

func NewUserUsecase(ur repository.IUserRepository, br repository.IBeanRepository, s3Service s3.IS3Service) IUserUsecase {
	return &userUsecase{ur, br, s3Service}
}
