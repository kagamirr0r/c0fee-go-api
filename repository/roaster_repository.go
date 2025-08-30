package repository

import (
	"c0fee-api/common"
	"c0fee-api/common/converter/entity_model"
	"c0fee-api/domain/entity"
	domainRepo "c0fee-api/domain/repository"
	"c0fee-api/model"

	"gorm.io/gorm"
)

type roasterRepository struct {
	db *gorm.DB
}

func (rr *roasterRepository) List(domainRoasters *[]entity.Roaster) error {
	var modelRoasters []model.Roaster
	if err := rr.db.Find(&modelRoasters).Error; err != nil {
		return err
	}

	// Convert model slice to domain entity slice
	*domainRoasters = entity_model.ModelRoastersToEntities(modelRoasters)
	return nil
}

func (rr *roasterRepository) Search(domainRoasters *[]entity.Roaster, params common.QueryParams) error {
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
	*domainRoasters = entity_model.ModelRoastersToEntities(modelRoasters)
	return nil
}

func NewRoasterRepository(db *gorm.DB) domainRepo.IRoasterRepository {
	return &roasterRepository{db}
}
