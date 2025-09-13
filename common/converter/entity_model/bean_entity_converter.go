package entity_model

import (
	"c0fee-api/domain/bean"
	"c0fee-api/domain/summary"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func BeanEntityToModel(beanEntity *bean.Entity) *model.Bean {
	if beanEntity == nil {
		return nil
	}

	modelBean := &model.Bean{
		ID:        beanEntity.ID,
		Name:      beanEntity.Name,
		Price:     beanEntity.Price,
		Currency:  model.Currency(beanEntity.Currency),
		ImageKey:  beanEntity.ImageKey,
		CreatedAt: beanEntity.CreatedAt,
		UpdatedAt: beanEntity.UpdatedAt,
	}

	// 必須フィールド
	modelBean.UserID = beanEntity.UserID
	// summary.User → model.User への変換
	modelBean.User = model.User{
		ID:   beanEntity.User.ID,
		Name: beanEntity.User.Name,
	}

	modelBean.RoasterID = beanEntity.RoasterID
	// summary.Roaster → model.Roaster への変換
	modelBean.Roaster = model.Roaster{
		ID:   beanEntity.Roaster.ID,
		Name: beanEntity.Roaster.Name,
	}

	modelBean.CountryID = beanEntity.CountryID
	// summary.Country → model.Country への変換
	modelBean.Country = model.Country{
		ID:   beanEntity.Country.ID,
		Name: beanEntity.Country.Name,
	}

	modelBean.RoastLevelID = beanEntity.RoastLevelID
	// summary.RoastLevel → model.RoastLevel への変換
	modelBean.RoastLevel = model.RoastLevel{
		ID:    beanEntity.RoastLevel.ID,
		Name:  beanEntity.RoastLevel.Name,
		Level: beanEntity.RoastLevel.Level,
	}

	// オプショナルフィールド
	if beanEntity.Area != nil && beanEntity.Area.ID != 0 {
		modelBean.AreaID = beanEntity.AreaID
		modelBean.Area = &model.Area{
			ID:        beanEntity.Area.ID,
			Name:      beanEntity.Area.Name,
			CountryID: beanEntity.Area.CountryID,
		}
	}

	if beanEntity.Farm != nil && beanEntity.Farm.ID != 0 {
		modelBean.FarmID = beanEntity.FarmID
		modelBean.Farm = &model.Farm{
			ID:     beanEntity.Farm.ID,
			Name:   beanEntity.Farm.Name,
			AreaID: beanEntity.Farm.AreaID,
		}
	}

	if beanEntity.Farmer != nil && beanEntity.Farmer.ID != 0 {
		modelBean.FarmerID = beanEntity.FarmerID
		modelBean.Farmer = &model.Farmer{
			ID:   beanEntity.Farmer.ID,
			Name: beanEntity.Farmer.Name,
		}
	}

	if beanEntity.ProcessMethod != nil && beanEntity.ProcessMethod.ID != 0 {
		modelBean.ProcessMethodID = beanEntity.ProcessMethodID
		modelBean.ProcessMethod = &model.ProcessMethod{
			ID:   beanEntity.ProcessMethod.ID,
			Name: beanEntity.ProcessMethod.Name,
		}
	}

	if len(beanEntity.Varieties) > 0 {
		modelBean.Varieties = make([]model.Variety, len(beanEntity.Varieties))
		for i, variety := range beanEntity.Varieties {
			modelBean.Varieties[i] = model.Variety{
				ID:   variety.ID,
				Name: variety.Name,
			}
		}
	}

	if len(beanEntity.BeanRatings) > 0 {
		modelBean.BeanRatings = make([]model.BeanRating, len(beanEntity.BeanRatings))
		for i, rating := range beanEntity.BeanRatings {
			modelBean.BeanRatings[i] = model.BeanRating{
				ID:         rating.ID,
				UserID:     rating.User.ID,
				Bitterness: rating.Bitterness,
				Acidity:    rating.Acidity,
				Body:       rating.Body,
				FlavorNote: rating.FlavorNote,
			}
		}
	}

	return modelBean
}

// DB Model → Domain Entity
func ModelToBeanEntity(modelBean *model.Bean) *bean.Entity {
	if modelBean == nil {
		return nil
	}

	beanEntity := &bean.Entity{
		ID:        modelBean.ID,
		Name:      modelBean.Name,
		Price:     modelBean.Price,
		Currency:  bean.Currency(modelBean.Currency),
		ImageKey:  modelBean.ImageKey,
		CreatedAt: modelBean.CreatedAt,
		UpdatedAt: modelBean.UpdatedAt,
	}

	// Convert related entities and set IDs
	beanEntity.UserID = modelBean.UserID
	// model.User → summary.User への変換
	beanEntity.User = summary.User{
		ID:   modelBean.User.ID,
		Name: modelBean.User.Name,
	}

	beanEntity.RoasterID = modelBean.RoasterID
	// model.Roaster → summary.Roaster への変換
	beanEntity.Roaster = summary.Roaster{
		ID:   modelBean.Roaster.ID,
		Name: modelBean.Roaster.Name,
	}

	beanEntity.CountryID = modelBean.CountryID
	// model.Country → summary.Country への変換
	beanEntity.Country = summary.Country{
		ID:   modelBean.Country.ID,
		Name: modelBean.Country.Name,
	}

	if modelBean.Area != nil && modelBean.Area.ID != 0 {
		beanEntity.AreaID = modelBean.AreaID
		beanEntity.Area = &summary.Area{
			ID:        modelBean.Area.ID,
			Name:      modelBean.Area.Name,
			CountryID: modelBean.Area.CountryID,
		}
	}

	if modelBean.Farm != nil && modelBean.Farm.ID != 0 {
		beanEntity.FarmID = modelBean.FarmID
		beanEntity.Farm = &summary.Farm{
			ID:     modelBean.Farm.ID,
			Name:   modelBean.Farm.Name,
			AreaID: modelBean.Farm.AreaID,
		}
	}

	if modelBean.Farmer != nil && modelBean.Farmer.ID != 0 {
		beanEntity.FarmerID = modelBean.FarmerID
		beanEntity.Farmer = &summary.Farmer{
			ID:   modelBean.Farmer.ID,
			Name: modelBean.Farmer.Name,
		}
	}

	if modelBean.ProcessMethod != nil && modelBean.ProcessMethod.ID != 0 {
		beanEntity.ProcessMethodID = modelBean.ProcessMethodID
		beanEntity.ProcessMethod = &summary.ProcessMethod{
			ID:   modelBean.ProcessMethod.ID,
			Name: modelBean.ProcessMethod.Name,
		}
	}

	if modelBean.RoastLevel.ID != 0 {
		beanEntity.RoastLevelID = modelBean.RoastLevelID
		beanEntity.RoastLevel = summary.RoastLevel{
			ID:    modelBean.RoastLevel.ID,
			Name:  modelBean.RoastLevel.Name,
			Level: modelBean.RoastLevel.Level,
		}
	}

	if len(modelBean.Varieties) > 0 {
		beanEntity.Varieties = make([]summary.Variety, len(modelBean.Varieties))
		for i, variety := range modelBean.Varieties {
			beanEntity.Varieties[i] = summary.Variety{
				ID:   variety.ID,
				Name: variety.Name,
			}
		}
	}

	if len(modelBean.BeanRatings) > 0 {
		beanEntity.BeanRatings = make([]summary.BeanRating, len(modelBean.BeanRatings))
		for i, rating := range modelBean.BeanRatings {
			beanEntity.BeanRatings[i] = summary.BeanRating{
				ID:     rating.ID,
				UserID: rating.UserID,
				User: summary.User{
					ID:   rating.User.ID,
					Name: rating.User.Name,
				},
				Bitterness: rating.Bitterness,
				Acidity:    rating.Acidity,
				Body:       rating.Body,
				FlavorNote: rating.FlavorNote,
			}
		}
	}

	return beanEntity
}

// Model slice → Bean Entity slice
func ModelsToBeanEntities(modelBeans []model.Bean) []bean.Entity {
	entities := make([]bean.Entity, len(modelBeans))
	for i, modelBean := range modelBeans {
		entities[i] = *ModelToBeanEntity(&modelBean)
	}
	return entities
}

// Bean Entity slice → Model slice
func BeanEntitiesToModels(beanEntities []bean.Entity) []model.Bean {
	models := make([]model.Bean, len(beanEntities))
	for i, beanEntity := range beanEntities {
		models[i] = *BeanEntityToModel(&beanEntity)
	}
	return models
}
