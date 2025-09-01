package entity_model

import (
	"c0fee-api/domain/entity"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func EntityRoasterToModel(entityRoaster *entity.Roaster) *model.Roaster {
	if entityRoaster == nil {
		return nil
	}

	return &model.Roaster{
		ID:        entityRoaster.ID,
		Name:      entityRoaster.Name,
		Address:   entityRoaster.Address,
		WebURL:    entityRoaster.WebURL,
		ImageKey:  entityRoaster.ImageKey,
		CreatedAt: entityRoaster.CreatedAt,
		UpdatedAt: entityRoaster.UpdatedAt,
	}
}

// DB Model → Domain Entity
func ModelRoasterToEntity(modelRoaster *model.Roaster) *entity.Roaster {
	if modelRoaster == nil {
		return nil
	}

	return &entity.Roaster{
		ID:        modelRoaster.ID,
		Name:      modelRoaster.Name,
		Address:   modelRoaster.Address,
		WebURL:    modelRoaster.WebURL,
		ImageKey:  modelRoaster.ImageKey,
		CreatedAt: modelRoaster.CreatedAt,
		UpdatedAt: modelRoaster.UpdatedAt,
	}
}

// Convert slice of models to entities
func ModelRoastersToEntities(modelRoasters []model.Roaster) []entity.Roaster {
	entities := make([]entity.Roaster, len(modelRoasters))
	for i, model := range modelRoasters {
		entities[i] = *ModelRoasterToEntity(&model)
	}
	return entities
}

// Convert slice of entities to models
func EntityRoastersToModels(entityRoasters []entity.Roaster) []model.Roaster {
	models := make([]model.Roaster, len(entityRoasters))
	for i, entity := range entityRoasters {
		models[i] = *EntityRoasterToModel(&entity)
	}
	return models
}
