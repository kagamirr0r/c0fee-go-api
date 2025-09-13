package entity_model

import (
	"c0fee-api/domain/country"
	"c0fee-api/domain/summary"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func CountryEntityToModel(countryEntity *country.Entity) *model.Country {
	if countryEntity == nil {
		return nil
	}

	modelCountry := &model.Country{
		ID:        countryEntity.ID,
		Name:      countryEntity.Name,
		Code:      countryEntity.Code,
		CreatedAt: countryEntity.CreatedAt,
		UpdatedAt: countryEntity.UpdatedAt,
	}

	// Convert related entities - avoiding circular reference
	if len(countryEntity.Areas) > 0 {
		modelCountry.Areas = make([]model.Area, len(countryEntity.Areas))
		for i, area := range countryEntity.Areas {
			modelCountry.Areas[i] = model.Area{
				ID:        area.ID,
				Name:      area.Name,
				CountryID: area.CountryID,
			}
		}
	}

	return modelCountry
}

// DB Model → Domain Entity
func ModelToCountryEntity(modelCountry *model.Country) *country.Entity {
	if modelCountry == nil {
		return nil
	}

	countryEntity := &country.Entity{
		ID:        modelCountry.ID,
		Name:      modelCountry.Name,
		Code:      modelCountry.Code,
		CreatedAt: modelCountry.CreatedAt,
		UpdatedAt: modelCountry.UpdatedAt,
	}

	// Convert related entities - avoiding circular reference
	if len(modelCountry.Areas) > 0 {
		countryEntity.Areas = make([]summary.Area, len(modelCountry.Areas))
		for i, area := range modelCountry.Areas {
			countryEntity.Areas[i] = summary.Area{
				ID:        area.ID,
				Name:      area.Name,
				CountryID: area.CountryID,
			}
		}
	}

	return countryEntity
}

// Model slice → Country Entity slice
func ModelsToCountryEntities(modelCountries []model.Country) []country.Entity {
	entities := make([]country.Entity, len(modelCountries))
	for i, model := range modelCountries {
		entities[i] = *ModelToCountryEntity(&model)
	}
	return entities
}

// Country Entity slice → Model slice
func CountryEntitiesToModels(countryEntities []country.Entity) []model.Country {
	models := make([]model.Country, len(countryEntities))
	for i, entity := range countryEntities {
		models[i] = *CountryEntityToModel(&entity)
	}
	return models
}
