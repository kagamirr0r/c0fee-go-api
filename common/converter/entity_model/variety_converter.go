package entity_model

import (
	"c0fee-api/domain/variety"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func VarietyEntityToModel(entityVariety *variety.Entity) *model.Variety {
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
func ModelToVarietyEntity(modelVariety *model.Variety) *variety.Entity {
	if modelVariety == nil {
		return nil
	}

	return &variety.Entity{
		ID:        modelVariety.ID,
		Name:      modelVariety.Name,
		CreatedAt: modelVariety.CreatedAt,
		UpdatedAt: modelVariety.UpdatedAt,
	}
}

// Convert slice of models to entities
func ModelsToVarietyEntities(modelVarieties []model.Variety) []variety.Entity {
	entities := make([]variety.Entity, len(modelVarieties))
	for i, model := range modelVarieties {
		entities[i] = *ModelToVarietyEntity(&model)
	}
	return entities
}

// Convert slice of entities to models
func VarietyEntitiesToModels(entityVarieties []variety.Entity) []model.Variety {
	models := make([]model.Variety, len(entityVarieties))
	for i, entity := range entityVarieties {
		models[i] = *VarietyEntityToModel(&entity)
	}
	return models
}
