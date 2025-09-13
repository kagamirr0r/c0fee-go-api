package entity_model

import (
	"c0fee-api/domain/farm"
	"c0fee-api/domain/summary"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func FarmEntityToModel(farmEntity *farm.Entity) *model.Farm {
	if farmEntity == nil {
		return nil
	}

	modelFarm := &model.Farm{
		ID:        farmEntity.ID,
		Name:      farmEntity.Name,
		AreaID:    farmEntity.AreaID,
		CreatedAt: farmEntity.CreatedAt,
		UpdatedAt: farmEntity.UpdatedAt,
	}

	// Convert related entities - avoiding circular reference
	if farmEntity.Area.ID != 0 {
		modelFarm.Area = model.Area{
			ID:        farmEntity.Area.ID,
			Name:      farmEntity.Area.Name,
			CountryID: farmEntity.Area.CountryID,
		}
	}

	if len(farmEntity.Farmers) > 0 {
		modelFarm.Farmers = make([]model.Farmer, len(farmEntity.Farmers))
		for i, farmer := range farmEntity.Farmers {
			modelFarm.Farmers[i] = model.Farmer{
				ID:   farmer.ID,
				Name: farmer.Name,
			}
		}
	}

	return modelFarm
}

// DB Model → Domain Entity
func ModelToFarmEntity(modelFarm *model.Farm) *farm.Entity {
	if modelFarm == nil {
		return nil
	}

	farmEntity := &farm.Entity{
		ID:        modelFarm.ID,
		Name:      modelFarm.Name,
		AreaID:    modelFarm.AreaID,
		CreatedAt: modelFarm.CreatedAt,
		UpdatedAt: modelFarm.UpdatedAt,
	}

	// Convert related entities - avoiding circular reference
	if modelFarm.Area.ID != 0 {
		farmEntity.Area = summary.Area{
			ID:        modelFarm.Area.ID,
			Name:      modelFarm.Area.Name,
			CountryID: modelFarm.Area.CountryID,
		}
	}

	if len(modelFarm.Farmers) > 0 {
		farmEntity.Farmers = make([]summary.Farmer, len(modelFarm.Farmers))
		for i, farmer := range modelFarm.Farmers {
			farmEntity.Farmers[i] = summary.Farmer{
				ID:   farmer.ID,
				Name: farmer.Name,
			}
		}
	}

	return farmEntity
}

// Model slice → Farm Entity slice
func ModelsToFarmEntities(modelFarms []model.Farm) []farm.Entity {
	entities := make([]farm.Entity, len(modelFarms))
	for i, model := range modelFarms {
		entities[i] = *ModelToFarmEntity(&model)
	}
	return entities
}

// Farm Entity slice → Model slice
func FarmEntitiesToModels(farmEntities []farm.Entity) []model.Farm {
	models := make([]model.Farm, len(farmEntities))
	for i, entity := range farmEntities {
		models[i] = *FarmEntityToModel(&entity)
	}
	return models
}
