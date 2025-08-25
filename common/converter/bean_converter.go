package converter

import (
	"c0fee-api/dto"
	"c0fee-api/model"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// DTO -> Model
func ConvertBeanInputToBean(userID string, data dto.BeanInput) (model.Bean, []uint) {
	fmt.Println("data in converter:", data)
	bean := model.Bean{
		Name:         data.Name,
		UserID:       uuid.MustParse(userID),
		CountryID:    data.CountryID,
		RoasterID:    data.RoasterID,
		RoastLevelID: data.RoastLevelID,
		Price:        data.Price,
		Currency:     model.JPY, // デフォルト
	}

	// Optional fields
	if data.ID != nil {
		bean.ID = *data.ID
	}
	if data.AreaID != nil {
		bean.AreaID = data.AreaID
	}
	if data.FarmID != nil {
		bean.FarmID = data.FarmID
	}
	if data.FarmerID != nil {
		bean.FarmerID = data.FarmerID
	}
	if data.ProcessMethodID != nil {
		bean.ProcessMethodID = data.ProcessMethodID
	}

	return bean, data.VarietyIDs
}

// Model -> DTO
func ConvertBeanToOutput(bean *model.Bean, imageURL string) dto.BeanOutput {
	response := dto.BeanOutput{
		ID:        bean.ID,
		Name:      bean.Name,
		CreatedAt: bean.CreatedAt.Format(time.RFC3339),
		UpdatedAt: bean.UpdatedAt.Format(time.RFC3339),
	}

	if imageURL != "" {
		response.ImageURL = &imageURL
	}

	// User (IDは string)
	response.User = dto.IdNameSummary{
		ID:   bean.User.ID.String(),
		Name: bean.User.Name,
	}

	// Roaster, Country, RoastLevel etc. (IDは uint)
	response.Roaster = dto.IdNameSummary{ID: bean.Roaster.ID, Name: bean.Roaster.Name}
	response.Country = dto.IdNameSummary{ID: bean.Country.ID, Name: bean.Country.Name}
	response.RoastLevel = dto.IdNameSummary{ID: bean.RoastLevel.ID, Name: bean.RoastLevel.Name}

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
		response.Price = &dto.PriceSummary{
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
			User:       dto.IdNameSummary{ID: rating.User.ID.String(), Name: rating.User.Name},
			Bitterness: rating.Bitterness,
			Acidity:    rating.Acidity,
			Body:       rating.Body,
			FlavorNote: flavorNote,
		}
	}
	response.BeanRatings = ratings

	return response
}

func ConvertBeanToBeanSummary(bean *model.Bean, imageURL string) dto.BeanSummary {
	response := dto.BeanSummary{
		ID:        bean.ID,
		Name:      bean.Name,
		CreatedAt: bean.CreatedAt.Format(time.RFC3339),
		UpdatedAt: bean.UpdatedAt.Format(time.RFC3339),
	}

	if imageURL != "" {
		response.ImageURL = &imageURL
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

	return response
}
