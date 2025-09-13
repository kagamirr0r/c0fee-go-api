package entity_model

import (
	"c0fee-api/domain/bean_rating"
	"c0fee-api/domain/summary"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func BeanRatingEntityToModel(beanRatingEntity *bean_rating.Entity) *model.BeanRating {
	if beanRatingEntity == nil {
		return nil
	}

	modelBeanRating := &model.BeanRating{
		ID:         beanRatingEntity.ID,
		BeanID:     beanRatingEntity.BeanID,
		UserID:     beanRatingEntity.UserID,
		Bitterness: beanRatingEntity.Bitterness,
		Acidity:    beanRatingEntity.Acidity,
		Body:       beanRatingEntity.Body,
		FlavorNote: beanRatingEntity.FlavorNote,
		CreatedAt:  beanRatingEntity.CreatedAt,
		UpdatedAt:  beanRatingEntity.UpdatedAt,
	}

	// Convert related entities
	if beanRatingEntity.User.ID.String() != "00000000-0000-0000-0000-000000000000" {
		modelBeanRating.User = model.User{
			ID:   beanRatingEntity.User.ID,
			Name: beanRatingEntity.User.Name,
		}
	}

	if beanRatingEntity.Bean.ID != 0 {
		modelBeanRating.Bean = model.Bean{
			ID:   beanRatingEntity.Bean.ID,
			Name: beanRatingEntity.Bean.Name,
		}
	}

	return modelBeanRating
}

// DB Model → Domain Entity
func ModelToBeanRatingEntity(modelBeanRating *model.BeanRating) *bean_rating.Entity {
	if modelBeanRating == nil {
		return nil
	}

	beanRatingEntity := &bean_rating.Entity{
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
		beanRatingEntity.User = summary.User{
			ID:   modelBeanRating.User.ID,
			Name: modelBeanRating.User.Name,
		}
	}

	if modelBeanRating.Bean.ID != 0 {
		beanRatingEntity.Bean = summary.Bean{
			ID:   modelBeanRating.Bean.ID,
			Name: modelBeanRating.Bean.Name,
		}
	}

	return beanRatingEntity
}

// Model slice → Bean Rating Entity slice
func ModelsToBeanRatingEntities(modelBeanRatings []model.BeanRating) []bean_rating.Entity {
	entities := make([]bean_rating.Entity, len(modelBeanRatings))
	for i, modelBeanRating := range modelBeanRatings {
		entities[i] = *ModelToBeanRatingEntity(&modelBeanRating)
	}
	return entities
}

// Bean Rating Entity slice → Model slice
func BeanRatingEntitiesToModels(beanRatingEntities []bean_rating.Entity) []model.BeanRating {
	models := make([]model.BeanRating, len(beanRatingEntities))
	for i, beanRatingEntity := range beanRatingEntities {
		models[i] = *BeanRatingEntityToModel(&beanRatingEntity)
	}
	return models
}
