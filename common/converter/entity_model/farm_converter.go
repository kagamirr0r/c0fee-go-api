package entity_model

import (
	"c0fee-api/domain/entity"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func EntityFarmToModel(entityFarm *entity.Farm) *model.Farm {
	if entityFarm == nil {
		return nil
	}

	modelFarm := &model.Farm{
		ID:        entityFarm.ID,
		Name:      entityFarm.Name,
		AreaID:    entityFarm.AreaID,
		CreatedAt: entityFarm.CreatedAt,
		UpdatedAt: entityFarm.UpdatedAt,
	}

	// Convert related entities - avoiding circular reference
	if entityFarm.Area.ID != 0 {
		modelFarm.Area = model.Area{
			ID:        entityFarm.Area.ID,
			Name:      entityFarm.Area.Name,
			CountryID: entityFarm.Area.CountryID,
			CreatedAt: entityFarm.Area.CreatedAt,
			UpdatedAt: entityFarm.Area.UpdatedAt,
		}
	}

	if len(entityFarm.Farmers) > 0 {
		modelFarm.Farmers = make([]model.Farmer, len(entityFarm.Farmers))
		for i, farmer := range entityFarm.Farmers {
			modelFarm.Farmers[i] = model.Farmer{
				ID:        farmer.ID,
				Name:      farmer.Name,
				FarmID:    farmer.FarmID,
				CreatedAt: farmer.CreatedAt,
				UpdatedAt: farmer.UpdatedAt,
			}
		}
	}

	return modelFarm
}

// DB Model → Domain Entity
func ModelFarmToEntity(modelFarm *model.Farm) *entity.Farm {
	if modelFarm == nil {
		return nil
	}

	entityFarm := &entity.Farm{
		ID:        modelFarm.ID,
		Name:      modelFarm.Name,
		AreaID:    modelFarm.AreaID,
		CreatedAt: modelFarm.CreatedAt,
		UpdatedAt: modelFarm.UpdatedAt,
	}

	// Convert related entities - avoiding circular reference
	if modelFarm.Area.ID != 0 {
		entityFarm.Area = entity.Area{
			ID:        modelFarm.Area.ID,
			Name:      modelFarm.Area.Name,
			CountryID: modelFarm.Area.CountryID,
			CreatedAt: modelFarm.Area.CreatedAt,
			UpdatedAt: modelFarm.Area.UpdatedAt,
		}
	}

	if len(modelFarm.Farmers) > 0 {
		entityFarm.Farmers = make([]entity.Farmer, len(modelFarm.Farmers))
		for i, farmer := range modelFarm.Farmers {
			entityFarm.Farmers[i] = entity.Farmer{
				ID:        farmer.ID,
				Name:      farmer.Name,
				FarmID:    farmer.FarmID,
				CreatedAt: farmer.CreatedAt,
				UpdatedAt: farmer.UpdatedAt,
			}
		}
	}

	return entityFarm
}

// Convert slice of models to entities
func ModelFarmsToEntities(modelFarms []model.Farm) []entity.Farm {
	entities := make([]entity.Farm, len(modelFarms))
	for i, model := range modelFarms {
		entities[i] = *ModelFarmToEntity(&model)
	}
	return entities
}

// Convert slice of entities to models
func EntityFarmsToModels(entityFarms []entity.Farm) []model.Farm {
	models := make([]model.Farm, len(entityFarms))
	for i, entity := range entityFarms {
		models[i] = *EntityFarmToModel(&entity)
	}
	return models
}
