package entity_model

import (
	"c0fee-api/domain/entity"
	"c0fee-api/model"
)

// Domain Entity → DB Model
func EntityBeanToModel(entityBean *entity.Bean) *model.Bean {
	if entityBean == nil {
		return nil
	}

	modelBean := &model.Bean{
		ID:        entityBean.ID,
		Name:      entityBean.Name,
		Price:     entityBean.Price,
		Currency:  model.Currency(entityBean.Currency),
		ImageKey:  entityBean.ImageKey,
		CreatedAt: entityBean.CreatedAt,
		UpdatedAt: entityBean.UpdatedAt,
	}

	// 必須フィールド
	modelBean.UserID = entityBean.UserID
	modelBean.User = *EntityUserToModel(&entityBean.User)

	modelBean.RoasterID = entityBean.RoasterID
	modelBean.Roaster = *EntityRoasterToModel(&entityBean.Roaster)

	modelBean.CountryID = entityBean.CountryID
	modelBean.Country = *EntityCountryToModel(&entityBean.Country)

	modelBean.RoastLevelID = entityBean.RoastLevelID
	modelBean.RoastLevel = *EntityRoastLevelToModel(&entityBean.RoastLevel)

	// オプショナルフィールド
	if entityBean.Area != nil && entityBean.Area.ID != 0 {
		modelBean.AreaID = entityBean.AreaID
		modelBean.Area = EntityAreaToModel(entityBean.Area)
	}

	if entityBean.Farm != nil && entityBean.Farm.ID != 0 {
		modelBean.FarmID = entityBean.FarmID
		modelBean.Farm = EntityFarmToModel(entityBean.Farm)
	}

	if entityBean.Farmer != nil && entityBean.Farmer.ID != 0 {
		modelBean.FarmerID = entityBean.FarmerID
		modelBean.Farmer = EntityFarmerToModel(entityBean.Farmer)
	}

	if entityBean.ProcessMethod != nil && entityBean.ProcessMethod.ID != 0 {
		modelBean.ProcessMethodID = entityBean.ProcessMethodID
		modelBean.ProcessMethod = EntityProcessMethodToModel(entityBean.ProcessMethod)
	}

	if len(entityBean.Varieties) > 0 {
		modelBean.Varieties = EntityVarietiesToModels(entityBean.Varieties)
	}

	if len(entityBean.BeanRatings) > 0 {
		modelBean.BeanRatings = EntityBeanRatingsToModels(entityBean.BeanRatings)
	}

	return modelBean
}

// DB Model → Domain Entity
func ModelBeanToEntity(modelBean *model.Bean) *entity.Bean {
	if modelBean == nil {
		return nil
	}

	entityBean := &entity.Bean{
		ID:        modelBean.ID,
		Name:      modelBean.Name,
		Price:     modelBean.Price,
		Currency:  entity.Currency(modelBean.Currency),
		ImageKey:  modelBean.ImageKey,
		CreatedAt: modelBean.CreatedAt,
		UpdatedAt: modelBean.UpdatedAt,
	}

	// Convert related entities and set IDs
	entityBean.UserID = modelBean.UserID
	entityBean.User = *ModelUserToEntity(&modelBean.User)

	entityBean.RoasterID = modelBean.RoasterID
	entityBean.Roaster = *ModelRoasterToEntity(&modelBean.Roaster)

	entityBean.CountryID = modelBean.CountryID
	entityBean.Country = *ModelCountryToEntity(&modelBean.Country)

	if modelBean.Area != nil && modelBean.Area.ID != 0 {
		entityBean.AreaID = modelBean.AreaID
		entityBean.Area = ModelAreaToEntity(modelBean.Area)
	}

	if modelBean.Farm != nil && modelBean.Farm.ID != 0 {
		entityBean.FarmID = modelBean.FarmID
		entityBean.Farm = ModelFarmToEntity(modelBean.Farm)
	}

	if modelBean.Farmer != nil && modelBean.Farmer.ID != 0 {
		entityBean.FarmerID = modelBean.FarmerID
		entityBean.Farmer = ModelFarmerToEntity(modelBean.Farmer)
	}

	if modelBean.ProcessMethod != nil && modelBean.ProcessMethod.ID != 0 {
		entityBean.ProcessMethodID = modelBean.ProcessMethodID
		entityBean.ProcessMethod = ModelProcessMethodToEntity(modelBean.ProcessMethod)
	}

	if modelBean.RoastLevel.ID != 0 {
		entityBean.RoastLevelID = modelBean.RoastLevelID
		entityBean.RoastLevel = *ModelRoastLevelToEntity(&modelBean.RoastLevel)
	}

	if len(modelBean.Varieties) > 0 {
		entityBean.Varieties = ModelVarietiesToEntities(modelBean.Varieties)
	}

	if len(modelBean.BeanRatings) > 0 {
		entityBean.BeanRatings = ModelBeanRatingsToEntities(modelBean.BeanRatings)
	}

	return entityBean
}

// Model slice → Entity slice
func ModelBeansToEntities(modelBeans []model.Bean) []entity.Bean {
	entities := make([]entity.Bean, len(modelBeans))
	for i, modelBean := range modelBeans {
		entities[i] = *ModelBeanToEntity(&modelBean)
	}
	return entities
}

// Entity slice → Model slice
func EntityBeansToModels(entityBeans []entity.Bean) []model.Bean {
	models := make([]model.Bean, len(entityBeans))
	for i, entityBean := range entityBeans {
		models[i] = *EntityBeanToModel(&entityBean)
	}
	return models
}
