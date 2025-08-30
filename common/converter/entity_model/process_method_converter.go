package entity_model

import (
	"c0fee-api/domain/entity"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func EntityProcessMethodToModel(entityProcessMethod *entity.ProcessMethod) *model.ProcessMethod {
	if entityProcessMethod == nil {
		return nil
	}

	return &model.ProcessMethod{
		ID:        entityProcessMethod.ID,
		Name:      entityProcessMethod.Name,
		CreatedAt: entityProcessMethod.CreatedAt,
		UpdatedAt: entityProcessMethod.UpdatedAt,
	}
}

// DB Model → Domain Entity
func ModelProcessMethodToEntity(modelProcessMethod *model.ProcessMethod) *entity.ProcessMethod {
	if modelProcessMethod == nil {
		return nil
	}

	return &entity.ProcessMethod{
		ID:        modelProcessMethod.ID,
		Name:      modelProcessMethod.Name,
		CreatedAt: modelProcessMethod.CreatedAt,
		UpdatedAt: modelProcessMethod.UpdatedAt,
	}
}

// Convert slice of models to entities
func ModelProcessMethodsToEntities(modelProcessMethods []model.ProcessMethod) []entity.ProcessMethod {
	entities := make([]entity.ProcessMethod, len(modelProcessMethods))
	for i, model := range modelProcessMethods {
		entities[i] = *ModelProcessMethodToEntity(&model)
	}
	return entities
}

// Convert slice of entities to models
func EntityProcessMethodsToModels(entityProcessMethods []entity.ProcessMethod) []model.ProcessMethod {
	models := make([]model.ProcessMethod, len(entityProcessMethods))
	for i, entity := range entityProcessMethods {
		models[i] = *EntityProcessMethodToModel(&entity)
	}
	return models
}
