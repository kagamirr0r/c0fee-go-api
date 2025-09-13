package dto_entity

import (
	"c0fee-api/domain/bean"
	"c0fee-api/dto"
	"time"

	"github.com/google/uuid"
)

// DTO -> Entity
func DtoBeanToEntity(userID string, data dto.BeanInput) (bean.Entity, []uint) {
	beanEntity := bean.Entity{
		Name:         data.Name,
		UserID:       uuid.MustParse(userID),
		CountryID:    data.CountryID,
		RoasterID:    data.RoasterID,
		RoastLevelID: data.RoastLevelID,
		Price:        data.Price,
		Currency:     bean.JPY, // デフォルト
	}

	// Optional fields
	if data.ID != nil {
		beanEntity.ID = *data.ID
	}
	if data.AreaID != nil {
		beanEntity.AreaID = data.AreaID
	}
	if data.FarmID != nil {
		beanEntity.FarmID = data.FarmID
	}
	if data.FarmerID != nil {
		beanEntity.FarmerID = data.FarmerID
	}
	if data.ProcessMethodID != nil {
		beanEntity.ProcessMethodID = data.ProcessMethodID
	}

	return beanEntity, data.VarietyIDs
}

// Entity -> DTO
func EntityBeanToDto(beanEntity *bean.Entity, imageURL string) dto.BeanOutput {
	response := dto.BeanOutput{
		ID:        beanEntity.ID,
		Name:      beanEntity.Name,
		CreatedAt: beanEntity.CreatedAt.Format(time.RFC3339),
		UpdatedAt: beanEntity.UpdatedAt.Format(time.RFC3339),
	}

	if imageURL != "" {
		response.ImageURL = &imageURL
	}

	// User (IDは string)
	if beanEntity.User.ID.String() != "00000000-0000-0000-0000-000000000000" {
		response.User = dto.IdNameSummary{
			ID:   beanEntity.User.ID.String(),
			Name: beanEntity.User.Name,
		}
	} else {
		response.User = dto.IdNameSummary{
			ID:   nil,
			Name: "",
		}
	}

	// Roaster, Country, RoastLevel etc. (IDは uint)
	response.Roaster = dto.IdNameSummary{ID: beanEntity.Roaster.ID, Name: beanEntity.Roaster.Name}
	response.Country = dto.IdNameSummary{ID: beanEntity.Country.ID, Name: beanEntity.Country.Name}
	response.RoastLevel = dto.IdNameSummary{ID: beanEntity.RoastLevel.ID, Name: beanEntity.RoastLevel.Name}

	// Optional fields
	if beanEntity.Area != nil {
		response.Area = &dto.IdNameSummary{ID: beanEntity.Area.ID, Name: beanEntity.Area.Name}
	}
	if beanEntity.Farm != nil {
		response.Farm = &dto.IdNameSummary{ID: beanEntity.Farm.ID, Name: beanEntity.Farm.Name}
	}
	if beanEntity.Farmer != nil {
		response.Farmer = &dto.IdNameSummary{ID: beanEntity.Farmer.ID, Name: beanEntity.Farmer.Name}
	}
	if beanEntity.ProcessMethod != nil {
		response.ProcessMethod = &dto.IdNameSummary{ID: beanEntity.ProcessMethod.ID, Name: beanEntity.ProcessMethod.Name}
	}

	// Varieties
	varieties := make([]dto.IdNameSummary, len(beanEntity.Varieties))
	for i, variety := range beanEntity.Varieties {
		varieties[i] = dto.IdNameSummary{ID: variety.ID, Name: variety.Name}
	}
	response.Varieties = varieties

	// Price
	if beanEntity.Price != nil {
		response.Price = &dto.PriceSummary{
			Amount:   float64(*beanEntity.Price),
			Currency: string(beanEntity.Currency),
		}
	}

	// BeanRatings
	ratings := make([]dto.BeanRatingSummary, len(beanEntity.BeanRatings))
	for i, rating := range beanEntity.BeanRatings {
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

func EntityBeanToBeanSummary(beanEntity *bean.Entity, imageURL string) dto.BeanSummary {
	response := dto.BeanSummary{
		ID:        beanEntity.ID,
		Name:      beanEntity.Name,
		CreatedAt: beanEntity.CreatedAt.Format(time.RFC3339),
		UpdatedAt: beanEntity.UpdatedAt.Format(time.RFC3339),
	}

	if imageURL != "" {
		response.ImageURL = &imageURL
	}

	response.User = dto.IdNameSummary{
		ID:   beanEntity.User.ID.String(),
		Name: beanEntity.User.Name,
	}

	// Roaster, Country, etc. (IDは uint)
	response.Roaster = dto.IdNameSummary{ID: beanEntity.Roaster.ID, Name: beanEntity.Roaster.Name}
	response.Country = dto.IdNameSummary{ID: beanEntity.Country.ID, Name: beanEntity.Country.Name}

	// Optional fields
	if beanEntity.Area != nil {
		response.Area = &dto.IdNameSummary{ID: beanEntity.Area.ID, Name: beanEntity.Area.Name}
	}
	if beanEntity.Farm != nil {
		response.Farm = &dto.IdNameSummary{ID: beanEntity.Farm.ID, Name: beanEntity.Farm.Name}
	}
	if beanEntity.Farmer != nil {
		response.Farmer = &dto.IdNameSummary{ID: beanEntity.Farmer.ID, Name: beanEntity.Farmer.Name}
	}
	if beanEntity.ProcessMethod != nil {
		response.ProcessMethod = &dto.IdNameSummary{ID: beanEntity.ProcessMethod.ID, Name: beanEntity.ProcessMethod.Name}
	}

	// Varieties
	varieties := make([]dto.IdNameSummary, len(beanEntity.Varieties))
	for i, variety := range beanEntity.Varieties {
		varieties[i] = dto.IdNameSummary{ID: variety.ID, Name: variety.Name}
	}
	response.Varieties = varieties

	return response
}
