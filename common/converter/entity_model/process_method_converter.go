package entity_model

import (
	"c0fee-api/domain/process_method"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func ProcessMethodEntityToModel(processMethodEntity *process_method.Entity) *model.ProcessMethod {
	if processMethodEntity == nil {
		return nil
	}

	return &model.ProcessMethod{
		ID:        processMethodEntity.ID,
		Name:      processMethodEntity.Name,
		CreatedAt: processMethodEntity.CreatedAt,
		UpdatedAt: processMethodEntity.UpdatedAt,
	}
}

// DB Model → Domain Entity
func ModelToProcessMethodEntity(modelProcessMethod *model.ProcessMethod) *process_method.Entity {
	if modelProcessMethod == nil {
		return nil
	}

	return &process_method.Entity{
		ID:        modelProcessMethod.ID,
		Name:      modelProcessMethod.Name,
		CreatedAt: modelProcessMethod.CreatedAt,
		UpdatedAt: modelProcessMethod.UpdatedAt,
	}
}

// Model slice → Process Method Entity slice
func ModelsToProcessMethodEntities(modelProcessMethods []model.ProcessMethod) []process_method.Entity {
	entities := make([]process_method.Entity, len(modelProcessMethods))
	for i, model := range modelProcessMethods {
		entities[i] = *ModelToProcessMethodEntity(&model)
	}
	return entities
}

// Process Method Entity slice → Model slice
func ProcessMethodEntitiesToModels(processMethodEntities []process_method.Entity) []model.ProcessMethod {
	models := make([]model.ProcessMethod, len(processMethodEntities))
	for i, entity := range processMethodEntities {
		models[i] = *ProcessMethodEntityToModel(&entity)
	}
	return models
}
