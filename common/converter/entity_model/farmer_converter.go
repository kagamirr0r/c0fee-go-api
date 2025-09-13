package entity_model

import (
	"c0fee-api/domain/farmer"
	"c0fee-api/domain/summary"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func FarmerEntityToModel(farmerEntity *farmer.Entity) *model.Farmer {
	if farmerEntity == nil {
		return nil
	}

	modelFarmer := &model.Farmer{
		ID:        farmerEntity.ID,
		Name:      farmerEntity.Name,
		FarmID:    farmerEntity.FarmID,
		CreatedAt: farmerEntity.CreatedAt,
		UpdatedAt: farmerEntity.UpdatedAt,
	}

	// Convert related entities - avoiding circular reference
	if farmerEntity.Farm.ID != 0 {
		modelFarmer.Farm = model.Farm{
			ID:     farmerEntity.Farm.ID,
			Name:   farmerEntity.Farm.Name,
			AreaID: farmerEntity.Farm.AreaID,
		}
	}

	return modelFarmer
}

// DB Model → Domain Entity
func ModelToFarmerEntity(modelFarmer *model.Farmer) *farmer.Entity {
	if modelFarmer == nil {
		return nil
	}

	farmerEntity := &farmer.Entity{
		ID:        modelFarmer.ID,
		Name:      modelFarmer.Name,
		FarmID:    modelFarmer.FarmID,
		CreatedAt: modelFarmer.CreatedAt,
		UpdatedAt: modelFarmer.UpdatedAt,
	}

	// Convert related entities - avoiding circular reference
	if modelFarmer.Farm.ID != 0 {
		farmerEntity.Farm = summary.Farm{
			ID:     modelFarmer.Farm.ID,
			Name:   modelFarmer.Farm.Name,
			AreaID: modelFarmer.Farm.AreaID,
		}
	}

	return farmerEntity
}

// Model slice → Farmer Entity slice
func ModelsToFarmerEntities(modelFarmers []model.Farmer) []farmer.Entity {
	entities := make([]farmer.Entity, len(modelFarmers))
	for i, model := range modelFarmers {
		entities[i] = *ModelToFarmerEntity(&model)
	}
	return entities
}

// Farmer Entity slice → Model slice
func FarmerEntitiesToModels(farmerEntities []farmer.Entity) []model.Farmer {
	models := make([]model.Farmer, len(farmerEntities))
	for i, entity := range farmerEntities {
		models[i] = *FarmerEntityToModel(&entity)
	}
	return models
}
