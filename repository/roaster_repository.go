package repository

import (
	"c0fee-api/common"
	"c0fee-api/model"

	"gorm.io/gorm"
)

type IRoasterRepository interface {
	List(countries *[]model.Roaster) error
	Search(roasters *[]model.Roaster, params common.QueryParams) error
}

type roasterRepository struct {
	db *gorm.DB
}

func (rr *roasterRepository) List(roasters *[]model.Roaster) error {
	if err := rr.db.Find(roasters).Error; err != nil {
		return err
	}
	return nil
}

func (rr *roasterRepository) Search(roasters *[]model.Roaster, params common.QueryParams) error {
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
	if err := query.Find(roasters).Error; err != nil {
		return err
	}
	return nil
}

func NewRoasterRepository(db *gorm.DB) IRoasterRepository {
	return &roasterRepository{db}
}
