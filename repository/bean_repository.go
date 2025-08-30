package repository

import (
	"c0fee-api/common"
	"c0fee-api/common/converter/entity_model"
	"c0fee-api/domain/entity"
	domainRepo "c0fee-api/domain/repository"
	"c0fee-api/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type beanRepository struct {
	db *gorm.DB
}

func (br *beanRepository) GetById(domainBean *entity.Bean, id uint) error {
	var modelBean model.Bean
	if err := br.db.
		Preload("User").
		Preload("Roaster").
		Preload("Country").
		Preload("RoastLevel").
		Preload("ProcessMethod").
		Preload("Varieties").
		Preload("Area").
		Preload("Farm").
		Preload("Farmer").
		Preload("BeanRatings").
		Preload("BeanRatings.User").
		Where("id = ?", id).
		First(&modelBean).Error; err != nil {
		return err
	}

	// Convert model to domain entity
	entityBean := entity_model.ModelBeanToEntity(&modelBean)
	*domainBean = *entityBean
	return nil
}

func (br *beanRepository) GetBeansByUserId(domainBeans *[]entity.Bean, userID uuid.UUID, params common.QueryParams) error {
	limit := 10 // デフォルトの取得件数
	if params.Limit > 0 {
		limit = params.Limit
	}

	var modelBeans []model.Bean
	if err := br.db.
		Preload("User").
		Preload("Roaster").
		Preload("Country").
		Preload("RoastLevel").
		Preload("ProcessMethod").
		Preload("Varieties").
		Preload("Area").
		Preload("Farm").
		Preload("Farmer").
		Where("user_id = ?", userID).
		Limit(limit).
		Order("created_at desc").
		Find(&modelBeans).Error; err != nil {
		return err
	}

	// Convert model slice to domain entity slice
	*domainBeans = entity_model.ModelBeansToEntities(modelBeans)
	return nil
}

func (br *beanRepository) SearchBeansByUserId(domainBeans *[]entity.Bean, userID uuid.UUID, params common.QueryParams) error {
	query := br.db.
		Preload("User").
		Preload("Roaster").
		Preload("Country").
		Preload("RoastLevel").
		Preload("ProcessMethod").
		Preload("Varieties").
		Preload("Area").
		Preload("Farm").
		Preload("Farmer").
		Where("beans.user_id = ?", userID)

	// name_likeパラメータが存在する場合、関係テーブルを検索
	if params.NameLike != "" {
		searchTerm := "%" + params.NameLike + "%"
		query = query.Where(`
			EXISTS (SELECT 1 FROM countries WHERE countries.id = beans.country_id AND LOWER(countries.name) LIKE LOWER(?)) OR
			EXISTS (SELECT 1 FROM areas WHERE areas.id = beans.area_id AND LOWER(areas.name) LIKE LOWER(?)) OR
			EXISTS (SELECT 1 FROM farms WHERE farms.id = beans.farm_id AND LOWER(farms.name) LIKE LOWER(?)) OR
			EXISTS (SELECT 1 FROM farmers WHERE farmers.id = beans.farmer_id AND LOWER(farmers.name) LIKE LOWER(?)) OR
			EXISTS (SELECT 1 FROM process_methods WHERE process_methods.id = beans.process_method_id AND LOWER(process_methods.name) LIKE LOWER(?)) OR
			EXISTS (SELECT 1 FROM bean_varieties bv JOIN varieties v ON v.id = bv.variety_id WHERE bv.bean_id = beans.id AND LOWER(v.name) LIKE LOWER(?)) OR
			LOWER(beans.name) LIKE LOWER(?)
		`, searchTerm, searchTerm, searchTerm, searchTerm, searchTerm, searchTerm, searchTerm)
	}

	// カーソルページネーション
	if params.Cursor > 0 {
		// Cursorが指定された場合、IDが指定されたカーソル値より小さいものを取得（降順での「次のページ」）
		query = query.Where("beans.id < ?", params.Cursor)
	}

	// limit
	if params.Limit > 0 {
		query = query.Limit(params.Limit)
	}

	// Order
	query = query.Order("beans.id desc")

	// 実行
	var modelBeans []model.Bean
	if err := query.Find(&modelBeans).Error; err != nil {
		return err
	}

	// Convert model slice to domain entity slice
	*domainBeans = entity_model.ModelBeansToEntities(modelBeans)
	return nil
}

func (br *beanRepository) Create(domainBean *entity.Bean) error {
	modelBean := entity_model.EntityBeanToModel(domainBean)
	if err := br.db.Create(modelBean).Error; err != nil {
		return err
	}

	// Update domain entity with database-generated fields
	domainBean.ID = modelBean.ID
	domainBean.CreatedAt = modelBean.CreatedAt
	domainBean.UpdatedAt = modelBean.UpdatedAt

	return nil
}

func (br *beanRepository) Update(domainBean *entity.Bean) error {
	modelBean := entity_model.EntityBeanToModel(domainBean)
	if err := br.db.Save(modelBean).Error; err != nil {
		return err
	}

	// Update domain entity with database fields
	domainBean.UpdatedAt = modelBean.UpdatedAt

	return nil
}

func (br *beanRepository) SetVarieties(beanID uint, varietyIDs []uint) error {
	var bean model.Bean
	if err := br.db.First(&bean, beanID).Error; err != nil {
		return err
	}

	// 既存の品種リレーションを削除
	if err := br.db.Model(&bean).Association("Varieties").Clear(); err != nil {
		return err
	}

	if len(varietyIDs) > 0 {
		var varieties []model.Variety
		if err := br.db.Where("id IN ?", varietyIDs).Find(&varieties).Error; err != nil {
			return err
		}

		// 品種の関連付け
		if err := br.db.Model(&bean).Association("Varieties").Append(varieties); err != nil {
			return err
		}
	}

	return nil
}

func NewBeanRepository(db *gorm.DB) domainRepo.IBeanRepository {
	return &beanRepository{db}
}
