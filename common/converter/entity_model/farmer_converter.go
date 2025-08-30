package entity_model

import (
	"c0fee-api/domain/entity"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func EntityFarmerToModel(entityFarmer *entity.Farmer) *model.Farmer {
	if entityFarmer == nil {
		return nil
	}

	modelFarmer := &model.Farmer{
		ID:        entityFarmer.ID,
		Name:      entityFarmer.Name,
		FarmID:    entityFarmer.FarmID,
		CreatedAt: entityFarmer.CreatedAt,
		UpdatedAt: entityFarmer.UpdatedAt,
	}

	// Convert related entities - avoiding circular reference
	if entityFarmer.Farm.ID != 0 {
		modelFarmer.Farm = model.Farm{
			ID:        entityFarmer.Farm.ID,
			Name:      entityFarmer.Farm.Name,
			AreaID:    entityFarmer.Farm.AreaID,
			CreatedAt: entityFarmer.Farm.CreatedAt,
			UpdatedAt: entityFarmer.Farm.UpdatedAt,
		}
	}

	return modelFarmer
}

// DB Model → Domain Entity
func ModelFarmerToEntity(modelFarmer *model.Farmer) *entity.Farmer {
	if modelFarmer == nil {
		return nil
	}

	entityFarmer := &entity.Farmer{
		ID:        modelFarmer.ID,
		Name:      modelFarmer.Name,
		FarmID:    modelFarmer.FarmID,
		CreatedAt: modelFarmer.CreatedAt,
		UpdatedAt: modelFarmer.UpdatedAt,
	}

	// Convert related entities - avoiding circular reference
	if modelFarmer.Farm.ID != 0 {
		entityFarmer.Farm = entity.Farm{
			ID:        modelFarmer.Farm.ID,
			Name:      modelFarmer.Farm.Name,
			AreaID:    modelFarmer.Farm.AreaID,
			CreatedAt: modelFarmer.Farm.CreatedAt,
			UpdatedAt: modelFarmer.Farm.UpdatedAt,
		}
	}

	return entityFarmer
}

// Convert slice of models to entities
func ModelFarmersToEntities(modelFarmers []model.Farmer) []entity.Farmer {
	entities := make([]entity.Farmer, len(modelFarmers))
	for i, model := range modelFarmers {
		entities[i] = *ModelFarmerToEntity(&model)
	}
	return entities
}

// Convert slice of entities to models
func EntityFarmersToModels(entityFarmers []entity.Farmer) []model.Farmer {
	models := make([]model.Farmer, len(entityFarmers))
	for i, entity := range entityFarmers {
		models[i] = *EntityFarmerToModel(&entity)
	}
	return models
}
