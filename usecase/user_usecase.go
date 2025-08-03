package usecase

import (
	"c0fee-api/common"
	"c0fee-api/dto"
	"c0fee-api/infrastructure/s3"
	"c0fee-api/model"
	"c0fee-api/repository"
	"fmt"
	"time"
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
		beansResponse[i] = uu.convertToBeanResponse(&bean, imageURL)
	}

	return dto.UserBeansOutput{
		User:  userResponse,
		Beans: beansResponse,
		Count: uint(len(beans)),
	}, nil
}

func (uu *userUsecase) convertToBeanResponse(bean *model.Bean, imageURL string) dto.BeanResponse {
	response := dto.BeanResponse{
		ID:         bean.ID,
		Name:       bean.Name,
		RoastLevel: string(bean.RoastLevel),
		CreatedAt:  bean.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  bean.UpdatedAt.Format(time.RFC3339),
	}

	if imageURL != "" {
		response.ImageURL = &imageURL
	}

	// User
	response.User = dto.UserSummary{
		ID:   bean.User.ID.String(),
		Name: bean.User.Name,
	}

	// Roaster
	response.Roaster = dto.RoasterSummary{
		ID:   bean.Roaster.ID,
		Name: bean.Roaster.Name,
	}

	// Country
	response.Country = dto.CountrySummary{
		ID:   bean.Country.ID,
		Name: bean.Country.Name,
	}

	// Optional fields
	if bean.Area != nil {
		response.Area = &dto.AreaSummary{
			ID:   bean.Area.ID,
			Name: bean.Area.Name,
		}
	}

	if bean.Farm != nil {
		response.Farm = &dto.FarmSummary{
			ID:   bean.Farm.ID,
			Name: bean.Farm.Name,
		}
	}

	if bean.Farmer != nil {
		response.Farmer = &dto.FarmerSummary{
			ID:   bean.Farmer.ID,
			Name: bean.Farmer.Name,
		}
	}

	if bean.ProcessMethod != nil {
		response.ProcessMethod = &dto.ProcessMethodSummary{
			ID:   bean.ProcessMethod.ID,
			Name: bean.ProcessMethod.Name,
		}
	}

	// Varieties
	varieties := make([]dto.VarietySummary, len(bean.Varieties))
	for i, variety := range bean.Varieties {
		varieties[i] = dto.VarietySummary{
			ID:   variety.ID,
			Name: variety.Name,
		}
	}
	response.Varieties = varieties

	// Price
	if bean.Price != nil {
		response.Price = &dto.PriceResponse{
			Amount:         float64(*bean.Price),
			Currency:       string(bean.Currency),
			FormattedPrice: uu.formatPrice(*bean.Price, bean.Currency),
		}
	}

	// BeanRatings
	ratings := make([]dto.BeanRatingSummary, len(bean.BeanRatings))
	for i, rating := range bean.BeanRatings {
		ratings[i] = dto.BeanRatingSummary{
			ID:         rating.ID,
			Bitterness: rating.Bitterness,
			Acidity:    rating.Acidity,
			Body:       rating.Body,
			FlavorNote: &rating.FlavorNote,
		}
	}
	response.BeanRatings = ratings

	return response
}

func (uu *userUsecase) formatPrice(price uint, currency model.Currency) string {
	switch currency {
	case model.JPY:
		return fmt.Sprintf("¥%d", price)
	case model.USD:
		return fmt.Sprintf("$%.2f", float64(price)/100)
	case model.EUR:
		return fmt.Sprintf("€%.2f", float64(price)/100)
	case model.GBP:
		return fmt.Sprintf("£%.2f", float64(price)/100)
	case model.KRW:
		return fmt.Sprintf("₩%d", price)
	default:
		return fmt.Sprintf("%.2f %s", float64(price), currency)
	}
}

func NewUserUsecase(ur repository.IUserRepository, bu repository.IBeanRepository, s3Service s3.IS3Service) IUserUsecase {
	return &userUsecase{ur, bu, s3Service}
}
