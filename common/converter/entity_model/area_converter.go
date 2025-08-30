package entity_model

import (
	"c0fee-api/domain/entity"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func EntityAreaToModel(entityArea *entity.Area) *model.Area {
	if entityArea == nil {
		return nil
	}

	modelArea := &model.Area{
		ID:        entityArea.ID,
		CountryID: entityArea.CountryID,
		Name:      entityArea.Name,
		CreatedAt: entityArea.CreatedAt,
		UpdatedAt: entityArea.UpdatedAt,
	}

	// Convert related entities - avoiding circular reference
	if entityArea.Country.ID != 0 {
		modelArea.Country = model.Country{
			ID:        entityArea.Country.ID,
			Name:      entityArea.Country.Name,
			Code:      entityArea.Country.Code,
			CreatedAt: entityArea.Country.CreatedAt,
			UpdatedAt: entityArea.Country.UpdatedAt,
		}
	}

	if len(entityArea.Farms) > 0 {
		modelArea.Farms = make([]model.Farm, len(entityArea.Farms))
		for i, farm := range entityArea.Farms {
			modelArea.Farms[i] = model.Farm{
				ID:        farm.ID,
				Name:      farm.Name,
				AreaID:    farm.AreaID,
				CreatedAt: farm.CreatedAt,
				UpdatedAt: farm.UpdatedAt,
			}
		}
	}

	return modelArea
}

// DB Model → Domain Entity
func ModelAreaToEntity(modelArea *model.Area) *entity.Area {
	if modelArea == nil {
		return nil
	}

	entityArea := &entity.Area{
		ID:        modelArea.ID,
		CountryID: modelArea.CountryID,
		Name:      modelArea.Name,
		CreatedAt: modelArea.CreatedAt,
		UpdatedAt: modelArea.UpdatedAt,
	}

	// Convert related entities - avoiding circular reference
	if modelArea.Country.ID != 0 {
		entityArea.Country = entity.Country{
			ID:        modelArea.Country.ID,
			Name:      modelArea.Country.Name,
			Code:      modelArea.Country.Code,
			CreatedAt: modelArea.Country.CreatedAt,
			UpdatedAt: modelArea.Country.UpdatedAt,
		}
	}

	if len(modelArea.Farms) > 0 {
		entityArea.Farms = make([]entity.Farm, len(modelArea.Farms))
		for i, farm := range modelArea.Farms {
			entityArea.Farms[i] = entity.Farm{
				ID:        farm.ID,
				Name:      farm.Name,
				AreaID:    farm.AreaID,
				CreatedAt: farm.CreatedAt,
				UpdatedAt: farm.UpdatedAt,
			}
		}
	}

	return entityArea
}

// Convert slice of models to entities
func ModelAreasToEntities(modelAreas []model.Area) []entity.Area {
	entities := make([]entity.Area, len(modelAreas))
	for i, model := range modelAreas {
		entities[i] = *ModelAreaToEntity(&model)
	}
	return entities
}

// Convert slice of entities to models
func EntityAreasToModels(entityAreas []entity.Area) []model.Area {
	models := make([]model.Area, len(entityAreas))
	for i, entity := range entityAreas {
		models[i] = *EntityAreaToModel(&entity)
	}
	return models
}
