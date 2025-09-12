package entity_model

import (
	"c0fee-api/domain/area"
	"c0fee-api/domain/summary"
	"c0fee-api/model"
)

// Domain Area Entity → DB Model
func AreaEntityToModel(areaEntity *area.Entity) *model.Area {
	if areaEntity == nil {
		return nil
	}

	modelArea := &model.Area{
		ID:        areaEntity.ID,
		CountryID: areaEntity.CountryID,
		Name:      areaEntity.Name,
		CreatedAt: areaEntity.CreatedAt,
		UpdatedAt: areaEntity.UpdatedAt,
	}

	// Convert related entities - avoiding circular reference
	if areaEntity.Country.ID != 0 {
		modelArea.Country = model.Country{
			ID:   areaEntity.Country.ID,
			Name: areaEntity.Country.Name,
		}
	}

	if len(areaEntity.Farms) > 0 {
		modelArea.Farms = make([]model.Farm, len(areaEntity.Farms))
		for i, farm := range areaEntity.Farms {
			modelArea.Farms[i] = model.Farm{
				ID:     farm.ID,
				Name:   farm.Name,
				AreaID: farm.AreaID,
			}
		}
	}

	return modelArea
}

// DB Model → Domain Area Entity
func ModelToAreaEntity(modelArea *model.Area) *area.Entity {
	if modelArea == nil {
		return nil
	}

	areaEntity := &area.Entity{
		ID:        modelArea.ID,
		CountryID: modelArea.CountryID,
		Name:      modelArea.Name,
		CreatedAt: modelArea.CreatedAt,
		UpdatedAt: modelArea.UpdatedAt,
	}

	// Convert related entities - avoiding circular reference
	if modelArea.Country.ID != 0 {
		areaEntity.Country = summary.Country{
			ID:   modelArea.Country.ID,
			Name: modelArea.Country.Name,
		}
	}

	if len(modelArea.Farms) > 0 {
		areaEntity.Farms = make([]summary.Farm, len(modelArea.Farms))
		for i, farm := range modelArea.Farms {
			areaEntity.Farms[i] = summary.Farm{
				ID:     farm.ID,
				Name:   farm.Name,
				AreaID: farm.AreaID,
			}
		}
	}

	return areaEntity
}

// Convert slice of models to area entities
func ModelsToAreaEntities(modelAreas []model.Area) []area.Entity {
	entities := make([]area.Entity, len(modelAreas))
	for i, model := range modelAreas {
		entities[i] = *ModelToAreaEntity(&model)
	}
	return entities
}

// Convert slice of area entities to models
func AreaEntitiesToModels(areaEntities []area.Entity) []model.Area {
	models := make([]model.Area, len(areaEntities))
	for i, entity := range areaEntities {
		models[i] = *AreaEntityToModel(&entity)
	}
	return models
}
