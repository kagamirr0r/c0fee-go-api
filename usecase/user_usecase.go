package usecase

import (
	"c0fee-api/common"
	"c0fee-api/common/converter/dto_entity"
	"c0fee-api/domain/entity"
	domainRepo "c0fee-api/domain/repository"
	"c0fee-api/dto"
	"c0fee-api/infrastructure/s3"

	"github.com/google/uuid"
)

type IUserUsecase interface {
	Create(userData dto.UserInput) (dto.UserOutput, error)
	Read(userID uuid.UUID) (dto.UserOutput, error)
	GetUserBeans(userID uuid.UUID, params common.QueryParams) (dto.BeansOutput, error)
}

type userUsecase struct {
	ur        domainRepo.IUserRepository
	br        domainRepo.IBeanRepository
	s3Service s3.IS3Service
}

func (uu *userUsecase) Create(userData dto.UserInput) (dto.UserOutput, error) {
	// convert to Entity
	user := entity.User{ID: userData.ID, Name: userData.Name}

	if err := uu.ur.CreateUser(&user); err != nil {
		return dto.UserOutput{}, err
	}
	return dto.UserOutput{ID: user.ID.String(), Name: user.Name}, nil
}

func (uu *userUsecase) Read(userID uuid.UUID) (dto.UserOutput, error) {
	// convert to Entity
	user := entity.User{ID: userID}
	if err := uu.ur.GetById(&user, user.ID); err != nil {
		return dto.UserOutput{}, err
	}

	var avatarURL string
	if user.AvatarKey != "" {
		presignedURL, err := uu.s3Service.GenerateUserAvatarURL(user.AvatarKey)
		if err != nil {
			return dto.UserOutput{}, err
		}
		avatarURL = presignedURL
	}
	return dto.UserOutput{ID: user.ID.String(), Name: user.Name, AvatarURL: avatarURL}, nil
}

func (uu *userUsecase) GetUserBeans(userID uuid.UUID, params common.QueryParams) (dto.BeansOutput, error) {
	var domainUser entity.User
	if err := uu.ur.GetById(&domainUser, userID); err != nil {
		return dto.BeansOutput{}, err
	}

	var domainBeans []entity.Bean
	if params.NameLike != "" || params.Cursor > 0 {
		err := uu.br.SearchBeansByUserId(&domainBeans, domainUser.ID, params)
		if err != nil {
			return dto.BeansOutput{}, err
		}
	} else {
		err := uu.br.GetBeansByUserId(&domainBeans, domainUser.ID, params)
		if err != nil {
			return dto.BeansOutput{}, err
		}
	}

	userBeans := make([]dto.BeanSummary, len(domainBeans))
	for i, bean := range domainBeans {
		var imageURL string
		if bean.ImageKey != nil {
			url, err := uu.s3Service.GenerateBeanImageURL(*bean.ImageKey)
			if err != nil {
				return dto.BeansOutput{}, err
			}
			imageURL = url
		}
		// Convert domain entity to model for existing converter function
		userBeans[i] = dto_entity.EntityBeanToBeanSummary(&bean, imageURL)
	}

	// カーソルページネーション用の情報を生成
	var nextCursor *uint
	if len(domainBeans) > 0 && params.Limit > 0 && len(domainBeans) == params.Limit {
		// 最後のBeanのIDをnext_cursorとして設定
		lastBeanID := domainBeans[len(domainBeans)-1].ID
		nextCursor = &lastBeanID
	}

	beansResponse := dto.BeansOutput{
		Beans:      userBeans,
		Count:      uint(len(domainBeans)),
		NextCursor: nextCursor,
	}

	return beansResponse, nil
}

func NewUserUsecase(ur domainRepo.IUserRepository, br domainRepo.IBeanRepository, s3Service s3.IS3Service) IUserUsecase {
	return &userUsecase{ur, br, s3Service}
}
