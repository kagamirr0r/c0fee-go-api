package usecase

import (
	"c0fee-api/common"
	"c0fee-api/common/converter/dto_entity"
	"c0fee-api/domain/entity"
	domainRepo "c0fee-api/domain/repository"
	"c0fee-api/dto"
	"c0fee-api/infrastructure/s3"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"slices"
	"strings"

	"github.com/google/uuid"
)

type IRoasterUsecase interface {
	List(params common.QueryParams) (dto.RoastersOutput, error)
	GetById(id uint) (dto.RoasterOutput, error)
	Create(userID string, data dto.RoasterInput, imageFile *multipart.FileHeader) (dto.RoasterOutput, error)
}

type roasterUsecase struct {
	ur        domainRepo.IUserRepository
	rr        domainRepo.IRoasterRepository
	br        domainRepo.IBeanRepository
	s3Service s3.IS3Service
}

func (ru *roasterUsecase) List(params common.QueryParams) (dto.RoastersOutput, error) {
	var roasters []entity.Roaster

	// パラメータが存在する場合は検索を使用、そうでなければリスト全体を取得
	if params.NameLike != "" || params.Limit > 0 {
		err := ru.rr.Search(&roasters, params)
		if err != nil {
			return dto.RoastersOutput{}, err
		}
	} else {
		err := ru.rr.List(&roasters)
		if err != nil {
			return dto.RoastersOutput{}, err
		}
	}

	roastersResponse := make([]dto.RoasterOutput, len(roasters))
	for i, roaster := range roasters {
		var imageURL *string
		if roaster.ImageKey != nil && *roaster.ImageKey != "" {
			url, err := ru.s3Service.GenerateRoasterImageURL(*roaster.ImageKey)
			if err == nil && url != "" {
				imageURL = &url
			}
		}

		roastersResponse[i] = dto.RoasterOutput{
			ID:       roaster.ID,
			Name:     roaster.Name,
			Address:  roaster.Address,
			WebURL:   roaster.WebURL,
			ImageURL: imageURL,
		}
	}

	return dto.RoastersOutput{Roasters: roastersResponse, Count: uint(len(roasters))}, nil
}

func (ru *roasterUsecase) GetById(id uint) (dto.RoasterOutput, error) {
	var roaster entity.Roaster
	err := ru.rr.GetById(&roaster, id)
	if err != nil {
		return dto.RoasterOutput{}, err
	}

	var imageURL *string
	if roaster.ImageKey != nil && *roaster.ImageKey != "" {
		url, err := ru.s3Service.GenerateRoasterImageURL(*roaster.ImageKey)
		if err == nil && url != "" {
			imageURL = &url
		}
	}

	// Use the preloaded beans from roaster entity
	beansOutput, err := dto_entity.BeanEntitiesToBeansOutput(roaster.Beans, common.QueryParams{}, ru.s3Service)
	if err != nil {
		return dto.RoasterOutput{}, err
	}

	return dto.RoasterOutput{
		ID:       roaster.ID,
		Name:     roaster.Name,
		Address:  roaster.Address,
		WebURL:   roaster.WebURL,
		ImageURL: imageURL,
		Beans:    beansOutput,
	}, nil
}

func (ru *roasterUsecase) Create(userID string, data dto.RoasterInput, imageFile *multipart.FileHeader) (dto.RoasterOutput, error) {
	// 共通バリデーション
	err := ru.validateInputData(userID, imageFile)
	if err != nil {
		return dto.RoasterOutput{}, err
	}

	// Domain Roasterエンティティを作成
	domainRoaster := entity.Roaster{
		Name:    data.Name,
		Address: data.Address,
	}

	if data.WebURL != nil {
		domainRoaster.WebURL = *data.WebURL
	}

	// 最初にRoasterを保存（画像なしで）
	if err := ru.rr.Create(&domainRoaster); err != nil {
		return dto.RoasterOutput{}, fmt.Errorf("failed to create roaster: %w", err)
	}

	// 画像をS3にアップロード（画像ファイルがある場合のみ）
	if imageFile != nil {
		imageKey, err := ru.s3Service.UploadRoasterImage(domainRoaster.ID, imageFile)
		if err != nil {
			return dto.RoasterOutput{}, fmt.Errorf("failed to upload image: %w", err)
		}

		// 画像キーのみを更新
		domainRoaster.ImageKey = &imageKey

		// 画像キーを更新
		if err := ru.rr.Update(&domainRoaster); err != nil {
			return dto.RoasterOutput{}, fmt.Errorf("failed to update roaster with image key: %w", err)
		}
	}

	// 作成されたRoasterを取得（関連データ含む）
	var createdRoaster entity.Roaster
	if err := ru.rr.GetById(&createdRoaster, domainRoaster.ID); err != nil {
		return dto.RoasterOutput{}, fmt.Errorf("failed to get created roaster: %w", err)
	}

	return ru.generateRoasterOutput(&createdRoaster)
}

func (ru *roasterUsecase) validateImageFile(imageFile *multipart.FileHeader) error {
	// ファイルサイズチェック（例：10MB 制限）
	maxSize := int64(10 * 1024 * 1024) // 10MB
	if imageFile.Size > maxSize {
		return errors.New("image file size must be less than 10MB")
	}

	// ファイル拡張子チェック
	ext := strings.ToLower(filepath.Ext(imageFile.Filename))
	allowedExts := []string{".jpg", ".jpeg", ".png", ".webp"}

	if !slices.Contains(allowedExts, ext) {
		return errors.New("image file must be jpg, jpeg, png, or webp")
	}

	return nil
}

// validateInputData は共通のバリデーション処理を実行します
func (ru *roasterUsecase) validateInputData(userID string, imageFile *multipart.FileHeader) error {
	// ユーザーの存在確認
	var user entity.User
	if err := ru.ur.GetById(&user, uuid.MustParse(userID)); err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	// 画像ファイルのバリデーション（ファイルがある場合のみ）
	if imageFile != nil {
		if err := ru.validateImageFile(imageFile); err != nil {
			return fmt.Errorf("invalid image file: %w", err)
		}
	}

	return nil
}

// generateRoasterOutput は画像URLを生成してRoasterOutputを作成します
func (ru *roasterUsecase) generateRoasterOutput(domainRoaster *entity.Roaster) (dto.RoasterOutput, error) {
	var imageURL *string
	if domainRoaster.ImageKey != nil && *domainRoaster.ImageKey != "" {
		url, err := ru.s3Service.GenerateRoasterImageURL(*domainRoaster.ImageKey)
		if err == nil && url != "" {
			imageURL = &url
		}
	}

	// Use the preloaded beans from roaster entity
	beansOutput, err := dto_entity.BeanEntitiesToBeansOutput(domainRoaster.Beans, common.QueryParams{}, ru.s3Service)
	if err != nil {
		return dto.RoasterOutput{}, err
	}

	return dto.RoasterOutput{
		ID:       domainRoaster.ID,
		Name:     domainRoaster.Name,
		Address:  domainRoaster.Address,
		WebURL:   domainRoaster.WebURL,
		ImageURL: imageURL,
		Beans:    beansOutput,
	}, nil
}

func NewRoasterUsecase(ur domainRepo.IUserRepository, rr domainRepo.IRoasterRepository, br domainRepo.IBeanRepository, s3Service s3.IS3Service) IRoasterUsecase {
	return &roasterUsecase{ur, rr, br, s3Service}
}
