package converter

import (
	"c0fee-api/dto"
	"c0fee-api/model"
	"time"

	"github.com/google/uuid"
)

// DTO -> Model
func ConvertCreateBeanDataToBean(userID string, data dto.CreateBeanData) model.Bean {
	// RoastLevel の変換
	var roastLevel model.RoastLevelType
	switch data.RoastLevel {
	case 1:
		roastLevel = model.Light
	case 2:
		roastLevel = model.MediumLight
	case 3:
		roastLevel = model.Medium
	case 4:
		roastLevel = model.MediumDark
	case 5:
		roastLevel = model.Dark
	default:
		roastLevel = model.Medium
	}

	bean := model.Bean{
		Name:       data.Name,
		UserID:     uuid.MustParse(userID),
		CountryID:  data.Country.ID,
		RoasterID:  data.Roaster.ID,
		RoastLevel: roastLevel,
		Price:      data.Price,
		Currency:   model.JPY, // デフォルト
	}

	// Optional fields
	if data.Area != nil {
		bean.AreaID = &data.Area.ID
	}
	if data.Farm != nil {
		bean.FarmID = &data.Farm.ID
	}
	if data.ProcessMethod != nil {
		bean.ProcessMethodID = &data.ProcessMethod.ID
	}

	return bean
}

// Model -> DTO
func ConvertToBeanResponse(bean *model.Bean, imageURL string) dto.BeanOutput {
	response := dto.BeanOutput{
		ID:         bean.ID,
		Name:       bean.Name,
		RoastLevel: string(bean.RoastLevel),
		CreatedAt:  bean.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  bean.UpdatedAt.Format(time.RFC3339),
	}

	if imageURL != "" {
		response.ImageURL = &imageURL
	}

	// User (IDは string)
	response.User = dto.IdNameSummary{
		ID:   bean.User.ID.String(),
		Name: bean.User.Name,
	}

	// Roaster, Country, etc. (IDは uint)
	response.Roaster = dto.IdNameSummary{ID: bean.Roaster.ID, Name: bean.Roaster.Name}
	response.Country = dto.IdNameSummary{ID: bean.Country.ID, Name: bean.Country.Name}

	// Optional fields
	if bean.Area != nil {
		response.Area = &dto.IdNameSummary{ID: bean.Area.ID, Name: bean.Area.Name}
	}
	if bean.Farm != nil {
		response.Farm = &dto.IdNameSummary{ID: bean.Farm.ID, Name: bean.Farm.Name}
	}
	if bean.Farmer != nil {
		response.Farmer = &dto.IdNameSummary{ID: bean.Farmer.ID, Name: bean.Farmer.Name}
	}
	if bean.ProcessMethod != nil {
		response.ProcessMethod = &dto.IdNameSummary{ID: bean.ProcessMethod.ID, Name: bean.ProcessMethod.Name}
	}

	// Varieties
	varieties := make([]dto.IdNameSummary, len(bean.Varieties))
	for i, variety := range bean.Varieties {
		varieties[i] = dto.IdNameSummary{ID: variety.ID, Name: variety.Name}
	}
	response.Varieties = varieties

	// Price
	if bean.Price != nil {
		response.Price = &dto.PriceOutput{
			Amount:   float64(*bean.Price),
			Currency: string(bean.Currency),
		}
	}

	// BeanRatings
	ratings := make([]dto.BeanRatingSummary, len(bean.BeanRatings))
	for i, rating := range bean.BeanRatings {
		var flavorNote *string
		if rating.FlavorNote != "" {
			flavorNote = &rating.FlavorNote
		}
		ratings[i] = dto.BeanRatingSummary{
			ID:         rating.ID,
			Bitterness: rating.Bitterness,
			Acidity:    rating.Acidity,
			Body:       rating.Body,
			FlavorNote: flavorNote,
		}
	}
	response.BeanRatings = ratings

	return response
}
