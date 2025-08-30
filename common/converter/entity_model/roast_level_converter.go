package entity_model

import (
	"c0fee-api/domain/entity"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func EntityRoastLevelToModel(entityRoastLevel *entity.RoastLevel) *model.RoastLevel {
	if entityRoastLevel == nil {
		return nil
	}

	return &model.RoastLevel{
		ID:        entityRoastLevel.ID,
		Name:      entityRoastLevel.Name,
		Level:     entityRoastLevel.Level,
		CreatedAt: entityRoastLevel.CreatedAt,
		UpdatedAt: entityRoastLevel.UpdatedAt,
	}
}

// DB Model → Domain Entity
func ModelRoastLevelToEntity(modelRoastLevel *model.RoastLevel) *entity.RoastLevel {
	if modelRoastLevel == nil {
		return nil
	}

	return &entity.RoastLevel{
		ID:        modelRoastLevel.ID,
		Name:      modelRoastLevel.Name,
		Level:     modelRoastLevel.Level,
		CreatedAt: modelRoastLevel.CreatedAt,
		UpdatedAt: modelRoastLevel.UpdatedAt,
	}
}

// Convert slice of models to entities
func ModelRoastLevelsToEntities(modelRoastLevels []model.RoastLevel) []entity.RoastLevel {
	entities := make([]entity.RoastLevel, len(modelRoastLevels))
	for i, model := range modelRoastLevels {
		entities[i] = *ModelRoastLevelToEntity(&model)
	}
	return entities
}

// Convert slice of entities to models
func EntityRoastLevelsToModels(entityRoastLevels []entity.RoastLevel) []model.RoastLevel {
	models := make([]model.RoastLevel, len(entityRoastLevels))
	for i, entity := range entityRoastLevels {
		models[i] = *EntityRoastLevelToModel(&entity)
	}
	return models
}
