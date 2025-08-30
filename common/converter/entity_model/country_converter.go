package entity_model

import (
	"c0fee-api/domain/entity"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func EntityCountryToModel(entityCountry *entity.Country) *model.Country {
	if entityCountry == nil {
		return nil
	}

	modelCountry := &model.Country{
		ID:        entityCountry.ID,
		Name:      entityCountry.Name,
		Code:      entityCountry.Code,
		CreatedAt: entityCountry.CreatedAt,
		UpdatedAt: entityCountry.UpdatedAt,
	}

	// Convert related entities - avoiding circular reference
	if len(entityCountry.Areas) > 0 {
		modelCountry.Areas = make([]model.Area, len(entityCountry.Areas))
		for i, area := range entityCountry.Areas {
			modelCountry.Areas[i] = model.Area{
				ID:        area.ID,
				Name:      area.Name,
				CountryID: area.CountryID,
				CreatedAt: area.CreatedAt,
				UpdatedAt: area.UpdatedAt,
			}
		}
	}

	return modelCountry
}

// DB Model → Domain Entity
func ModelCountryToEntity(modelCountry *model.Country) *entity.Country {
	if modelCountry == nil {
		return nil
	}

	entityCountry := &entity.Country{
		ID:        modelCountry.ID,
		Name:      modelCountry.Name,
		Code:      modelCountry.Code,
		CreatedAt: modelCountry.CreatedAt,
		UpdatedAt: modelCountry.UpdatedAt,
	}

	// Convert related entities - avoiding circular reference
	if len(modelCountry.Areas) > 0 {
		entityCountry.Areas = make([]entity.Area, len(modelCountry.Areas))
		for i, area := range modelCountry.Areas {
			entityCountry.Areas[i] = entity.Area{
				ID:        area.ID,
				Name:      area.Name,
				CountryID: area.CountryID,
				CreatedAt: area.CreatedAt,
				UpdatedAt: area.UpdatedAt,
			}
		}
	}

	return entityCountry
}

// Convert slice of models to entities
func ModelCountriesToEntities(modelCountries []model.Country) []entity.Country {
	entities := make([]entity.Country, len(modelCountries))
	for i, model := range modelCountries {
		entities[i] = *ModelCountryToEntity(&model)
	}
	return entities
}

// Convert slice of entities to models
func EntityCountriesToModels(entityCountries []entity.Country) []model.Country {
	models := make([]model.Country, len(entityCountries))
	for i, entity := range entityCountries {
		models[i] = *EntityCountryToModel(&entity)
	}
	return models
}
