package entity_model

import (
	"c0fee-api/domain/entity"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func EntityBeanRatingToModel(entityBeanRating *entity.BeanRating) *model.BeanRating {
	if entityBeanRating == nil {
		return nil
	}
	
	return &model.BeanRating{
		ID:         entityBeanRating.ID,
		BeanID:     entityBeanRating.BeanID,
		UserID:     entityBeanRating.UserID,
		Bitterness: entityBeanRating.Bitterness,
		Acidity:    entityBeanRating.Acidity,
		Body:       entityBeanRating.Body,
		FlavorNote: entityBeanRating.FlavorNote,
		CreatedAt:  entityBeanRating.CreatedAt,
		UpdatedAt:  entityBeanRating.UpdatedAt,
	}
}

// DB Model → Domain Entity
func ModelBeanRatingToEntity(modelBeanRating *model.BeanRating) *entity.BeanRating {
	if modelBeanRating == nil {
		return nil
	}
	
	entityBeanRating := &entity.BeanRating{
		ID:         modelBeanRating.ID,
		BeanID:     modelBeanRating.BeanID,
		UserID:     modelBeanRating.UserID,
		Bitterness: modelBeanRating.Bitterness,
		Acidity:    modelBeanRating.Acidity,
		Body:       modelBeanRating.Body,
		FlavorNote: modelBeanRating.FlavorNote,
		CreatedAt:  modelBeanRating.CreatedAt,
		UpdatedAt:  modelBeanRating.UpdatedAt,
	}
	
	// Convert related entities
	if modelBeanRating.User.ID.String() != "00000000-0000-0000-0000-000000000000" {
		entityBeanRating.User = *ModelUserToEntity(&modelBeanRating.User)
	}
	
	return entityBeanRating
}

// Model slice → Entity slice
func ModelBeanRatingsToEntities(modelBeanRatings []model.BeanRating) []entity.BeanRating {
	entities := make([]entity.BeanRating, len(modelBeanRatings))
	for i, modelBeanRating := range modelBeanRatings {
		entities[i] = *ModelBeanRatingToEntity(&modelBeanRating)
	}
	return entities
}

// Entity slice → Model slice
func EntityBeanRatingsToModels(entityBeanRatings []entity.BeanRating) []model.BeanRating {
	models := make([]model.BeanRating, len(entityBeanRatings))
	for i, entityBeanRating := range entityBeanRatings {
		models[i] = *EntityBeanRatingToModel(&entityBeanRating)
	}
	return models
}