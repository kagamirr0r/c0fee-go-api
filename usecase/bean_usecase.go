package usecase

import (
	"c0fee-api/dto"
	"c0fee-api/infrastructure/s3"
	"c0fee-api/model"
	"c0fee-api/repository"
	"fmt"
	"time"
)

type IBeanUsecase interface {
	Read(bean model.Bean) (dto.BeanResponse, error)
}

type beanUsecase struct {
	ur        repository.IUserRepository
	br        repository.IBeanRepository
	s3Service s3.IS3Service
}

func (bu *beanUsecase) Read(bean model.Bean) (dto.BeanOutput, error) {
	storedBean := model.Bean{}
	if err := bu.br.GetById(&storedBean, bean.ID); err != nil {
		return dto.BeanOutput{}, err
	}

	var imageURL string
	if storedBean.ImageKey != nil {
		url, err := bu.s3Service.GenerateBeanImageURL(*storedBean.ImageKey)
		if err != nil {
			return dto.BeanOutput{}, err
		}
		imageURL = url
	}

	return converter.ConvertToBeanResponse(&storedBean, imageURL), nil
}

func (bu *beanUsecase) convertToBeanResponse(bean *model.Bean, imageURL string) dto.BeanResponse {
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
			FormattedPrice: bu.formatPrice(*bean.Price, bean.Currency),
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

func (bu *beanUsecase) formatPrice(price uint, currency model.Currency) string {
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

func NewBeanUsecase(ur repository.IUserRepository, br repository.IBeanRepository, s3Service s3.IS3Service) IBeanUsecase {
	return &beanUsecase{ur, br, s3Service}
}
