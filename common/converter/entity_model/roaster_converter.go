package entity_model

import (
	"c0fee-api/domain/roaster"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func RoasterEntityToModel(roasterEntity *roaster.Entity) *model.Roaster {
	if roasterEntity == nil {
		return nil
	}

	return &model.Roaster{
		ID:        roasterEntity.ID,
		Name:      roasterEntity.Name,
		Address:   roasterEntity.Address,
		WebURL:    roasterEntity.WebURL,
		ImageKey:  roasterEntity.ImageKey,
		CreatedAt: roasterEntity.CreatedAt,
		UpdatedAt: roasterEntity.UpdatedAt,
	}
}

// DB Model → Domain Entity
func ModelToRoasterEntity(modelRoaster *model.Roaster) *roaster.Entity {
	if modelRoaster == nil {
		return nil
	}

	return &roaster.Entity{
		ID:        modelRoaster.ID,
		Name:      modelRoaster.Name,
		Address:   modelRoaster.Address,
		WebURL:    modelRoaster.WebURL,
		ImageKey:  modelRoaster.ImageKey,
		CreatedAt: modelRoaster.CreatedAt,
		UpdatedAt: modelRoaster.UpdatedAt,
	}
}

// Model slice → Roaster Entity slice
func ModelsToRoasterEntities(modelRoasters []model.Roaster) []roaster.Entity {
	entities := make([]roaster.Entity, len(modelRoasters))
	for i, model := range modelRoasters {
		entities[i] = *ModelToRoasterEntity(&model)
	}
	return entities
}

// Roaster Entity slice → Model slice
func RoasterEntitiesToModels(roasterEntities []roaster.Entity) []model.Roaster {
	models := make([]model.Roaster, len(roasterEntities))
	for i, entity := range roasterEntities {
		models[i] = *RoasterEntityToModel(&entity)
	}
	return models
}
