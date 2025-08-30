package entity_model

import (
	"c0fee-api/domain/entity"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func EntityVarietyToModel(entityVariety *entity.Variety) *model.Variety {
	if entityVariety == nil {
		return nil
	}

	return &model.Variety{
		ID:        entityVariety.ID,
		Name:      entityVariety.Name,
		CreatedAt: entityVariety.CreatedAt,
		UpdatedAt: entityVariety.UpdatedAt,
	}
}

// DB Model → Domain Entity
func ModelVarietyToEntity(modelVariety *model.Variety) *entity.Variety {
	if modelVariety == nil {
		return nil
	}

	return &entity.Variety{
		ID:        modelVariety.ID,
		Name:      modelVariety.Name,
		CreatedAt: modelVariety.CreatedAt,
		UpdatedAt: modelVariety.UpdatedAt,
	}
}

// Convert slice of models to entities
func ModelVarietiesToEntities(modelVarieties []model.Variety) []entity.Variety {
	entities := make([]entity.Variety, len(modelVarieties))
	for i, model := range modelVarieties {
		entities[i] = *ModelVarietyToEntity(&model)
	}
	return entities
}

// Convert slice of entities to models
func EntityVarietiesToModels(entityVarieties []entity.Variety) []model.Variety {
	models := make([]model.Variety, len(entityVarieties))
	for i, entity := range entityVarieties {
		models[i] = *EntityVarietyToModel(&entity)
	}
	return models
}
