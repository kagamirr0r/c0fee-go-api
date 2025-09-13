package entity_model

import (
	"c0fee-api/domain/roast_level"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func RoastLevelEntityToModel(roastLevelEntity *roast_level.Entity) *model.RoastLevel {
	if roastLevelEntity == nil {
		return nil
	}

	return &model.RoastLevel{
		ID:        roastLevelEntity.ID,
		Name:      roastLevelEntity.Name,
		Level:     roastLevelEntity.Level,
		CreatedAt: roastLevelEntity.CreatedAt,
		UpdatedAt: roastLevelEntity.UpdatedAt,
	}
}

// DB Model → Domain Entity
func ModelToRoastLevelEntity(modelRoastLevel *model.RoastLevel) *roast_level.Entity {
	if modelRoastLevel == nil {
		return nil
	}

	return &roast_level.Entity{
		ID:        modelRoastLevel.ID,
		Name:      modelRoastLevel.Name,
		Level:     modelRoastLevel.Level,
		CreatedAt: modelRoastLevel.CreatedAt,
		UpdatedAt: modelRoastLevel.UpdatedAt,
	}
}

// Model slice → Roast Level Entity slice
func ModelsToRoastLevelEntities(modelRoastLevels []model.RoastLevel) []roast_level.Entity {
	entities := make([]roast_level.Entity, len(modelRoastLevels))
	for i, model := range modelRoastLevels {
		entities[i] = *ModelToRoastLevelEntity(&model)
	}
	return entities
}

// Roast Level Entity slice → Model slice
func RoastLevelEntitiesToModels(roastLevelEntities []roast_level.Entity) []model.RoastLevel {
	models := make([]model.RoastLevel, len(roastLevelEntities))
	for i, entity := range roastLevelEntities {
		models[i] = *RoastLevelEntityToModel(&entity)
	}
	return models
}
