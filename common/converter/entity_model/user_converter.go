package entity_model

import (
	"c0fee-api/domain/entity"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func EntityUserToModel(entityUser *entity.User) *model.User {
	if entityUser == nil {
		return nil
	}

	return &model.User{
		ID:        entityUser.ID,
		Name:      entityUser.Name,
		AvatarKey: entityUser.AvatarKey,
		CreatedAt: entityUser.CreatedAt,
		UpdatedAt: entityUser.UpdatedAt,
	}
}

// DB Model → Domain Entity
func ModelUserToEntity(modelUser *model.User) *entity.User {
	if modelUser == nil {
		return nil
	}

	return &entity.User{
		ID:        modelUser.ID,
		Name:      modelUser.Name,
		AvatarKey: modelUser.AvatarKey,
		CreatedAt: modelUser.CreatedAt,
		UpdatedAt: modelUser.UpdatedAt,
	}
}

// Entity slice → Model slice
func EntityUsersToModels(entityUsers []entity.User) []model.User {
	models := make([]model.User, len(entityUsers))
	for i, entityUser := range entityUsers {
		models[i] = *EntityUserToModel(&entityUser)
	}
	return models
}

// Model slice → Entity slice
func ModelUsersToEntities(modelUsers []model.User) []entity.User {
	entities := make([]entity.User, len(modelUsers))
	for i, modelUser := range modelUsers {
		entities[i] = *ModelUserToEntity(&modelUser)
	}
	return entities
}
