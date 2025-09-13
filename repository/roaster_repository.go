package repository

import (
	"c0fee-api/common"
	"c0fee-api/common/converter/entity_model"
	"c0fee-api/domain/roaster"
	"c0fee-api/model"

	"gorm.io/gorm"
)

type roasterRepository struct {
	db *gorm.DB
}

func (rr *roasterRepository) List(domainRoasters *[]roaster.Entity) error {
	var modelRoasters []model.Roaster
	if err := rr.db.Find(&modelRoasters).Error; err != nil {
		return err
	}

	// Convert model slice to domain entity slice
	*domainRoasters = entity_model.ModelsToRoasterEntities(modelRoasters)
	return nil
}

func (rr *roasterRepository) Search(domainRoasters *[]roaster.Entity, params common.QueryParams) error {
	// 基本のクエリを初期化
	query := rr.db

	// name_likeパラメータが存在する場合、LIKE検索を追加
	if params.NameLike != "" {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+params.NameLike+"%")
	}

	// limitパラメータが存在する場合、制限を追加
	if params.Limit > 0 {
		query = query.Limit(params.Limit)
	}

	// 最終的なクエリを実行
	var modelRoasters []model.Roaster
	if err := query.Find(&modelRoasters).Error; err != nil {
		return err
	}

	// Convert model slice to domain entity slice
	*domainRoasters = entity_model.ModelsToRoasterEntities(modelRoasters)
	return nil
}

func (rr *roasterRepository) GetById(domainRoaster *roaster.Entity, id uint) error {
	var modelRoaster model.Roaster
	if err := rr.db.
		Preload("Beans").
		Preload("Beans.User").
		Preload("Beans.Roaster").
		Preload("Beans.Country").
		Preload("Beans.Area").
		Preload("Beans.Farm").
		Preload("Beans.Farmer").
		Preload("Beans.RoastLevel").
		Preload("Beans.ProcessMethod").
		Preload("Beans.Varieties").
		Preload("Beans.BeanRatings").
		Preload("Beans.BeanRatings.User").
		Where("id = ?", id).
		First(&modelRoaster).
		Error; err != nil {
		return err
	}

	*domainRoaster = *entity_model.ModelToRoasterEntity(&modelRoaster)
	return nil
}

func (rr *roasterRepository) Create(domainRoaster *roaster.Entity) error {
	modelRoaster := entity_model.RoasterEntityToModel(domainRoaster)

	if err := rr.db.Create(modelRoaster).Error; err != nil {
		return err
	}

	*domainRoaster = *entity_model.ModelToRoasterEntity(modelRoaster)
	return nil
}

func (rr *roasterRepository) Update(domainRoaster *roaster.Entity) error {
	modelRoaster := entity_model.RoasterEntityToModel(domainRoaster)

	if err := rr.db.Save(modelRoaster).Error; err != nil {
		return err
	}

	*domainRoaster = *entity_model.ModelToRoasterEntity(modelRoaster)
	return nil
}

func NewRoasterRepository(db *gorm.DB) roaster.IRoasterRepository {
	return &roasterRepository{db}
}
