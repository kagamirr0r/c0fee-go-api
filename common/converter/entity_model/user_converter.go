package entity_model

import (
	"c0fee-api/domain/user"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func UserEntityToModel(userEntity *user.Entity) *model.User {
	if userEntity == nil {
		return nil
	}

	return &model.User{
		ID:        userEntity.ID,
		Name:      userEntity.Name,
		AvatarKey: userEntity.AvatarKey,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
	}
}

// DB Model → Domain Entity
func ModelToUserEntity(modelUser *model.User) *user.Entity {
	if modelUser == nil {
		return nil
	}

	return &user.Entity{
		ID:        modelUser.ID,
		Name:      modelUser.Name,
		AvatarKey: modelUser.AvatarKey,
		CreatedAt: modelUser.CreatedAt,
		UpdatedAt: modelUser.UpdatedAt,
	}
}

// User Entity slice → Model slice
func UserEntitiesToModels(userEntities []user.Entity) []model.User {
	models := make([]model.User, len(userEntities))
	for i, userEntity := range userEntities {
		models[i] = *UserEntityToModel(&userEntity)
	}
	return models
}

// Model slice → User Entity slice
func ModelsToUserEntities(modelUsers []model.User) []user.Entity {
	entities := make([]user.Entity, len(modelUsers))
	for i, modelUser := range modelUsers {
		entities[i] = *ModelToUserEntity(&modelUser)
	}
	return entities
}
