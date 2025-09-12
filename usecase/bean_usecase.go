package usecase

import (
	"c0fee-api/common/converter/dto_entity"
	"c0fee-api/domain/bean"
	"c0fee-api/domain/bean_rating"
	"c0fee-api/domain/user"
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

type IBeanUsecase interface {
	Read(beanID uint) (dto.BeanOutput, error)
	Create(userID string, data dto.BeanInput, imageFile *multipart.FileHeader) (dto.BeanOutput, error)
	Update(beanID uint, userID string, data dto.BeanInput, imageFile *multipart.FileHeader) (dto.BeanOutput, error)
}

type beanUsecase struct {
	ur        user.IUserRepository
	br        bean.IBeanRepository
	brr       bean_rating.IBeanRatingRepository
	s3Service s3.IS3Service
}

func (bu *beanUsecase) Read(beanID uint) (dto.BeanOutput, error) {
	var domainBean bean.Entity
	if err := bu.br.GetById(&domainBean, beanID); err != nil {
		return dto.BeanOutput{}, err
	}

	var imageURL string
	if domainBean.ImageKey != nil {
		url, err := bu.s3Service.GenerateBeanImageURL(*domainBean.ImageKey)
		if err != nil {
			return dto.BeanOutput{}, err
		}
		imageURL = url
	}

	// Convert domain entity directly to DTO
	return dto_entity.EntityBeanToDto(&domainBean, imageURL), nil
}

func (bu *beanUsecase) Create(userID string, data dto.BeanInput, imageFile *multipart.FileHeader) (dto.BeanOutput, error) {
	// 共通バリデーション
	err := bu.validateInputData(userID, imageFile)
	if err != nil {
		return dto.BeanOutput{}, err
	}

	// Domain Beanエンティティを作成
	domainBean, varietyIDs := dto_entity.DtoBeanToEntity(userID, data)

	// 最初にBeanを保存（画像なしで）
	if err := bu.br.Create(&domainBean); err != nil {
		return dto.BeanOutput{}, fmt.Errorf("failed to create bean: %w", err)
	}

	// 品種の関連付け
	if len(varietyIDs) > 0 {
		if err := bu.br.SetVarieties(domainBean.ID, varietyIDs); err != nil {
			return dto.BeanOutput{}, fmt.Errorf("failed to set varieties: %w", err)
		}
	}

	// 画像をS3にアップロード（画像ファイルがある場合のみ）
	if imageFile != nil {
		imageKey, err := bu.s3Service.UploadBeanImage(domainBean.ID, imageFile)
		if err != nil {
			return dto.BeanOutput{}, fmt.Errorf("failed to upload image: %w", err)
		}

		// 画像キーのみを更新（domainBeanの他のフィールドはそのまま保持）
		domainBean.ImageKey = &imageKey

		// 画像キーを更新
		if err := bu.br.Update(&domainBean); err != nil {
			return dto.BeanOutput{}, fmt.Errorf("failed to update bean with image key: %w", err)
		}
	}

	// 作成されたBeanを取得（関連データ含む）
	var createdBean bean.Entity
	if err := bu.br.GetById(&createdBean, domainBean.ID); err != nil {
		return dto.BeanOutput{}, fmt.Errorf("failed to get created bean: %w", err)
	}

	// BeanRatingがある場合は作成
	if data.BeanRating != nil {
		if err := bu.createBeanRating(domainBean.ID, userID, data.BeanRating); err != nil {
			return dto.BeanOutput{}, err
		}

		// BeanRatingを含めて再取得
		if err := bu.br.GetById(&createdBean, domainBean.ID); err != nil {
			return dto.BeanOutput{}, fmt.Errorf("failed to get created bean with rating: %w", err)
		}
	}

	return bu.generateBeanOutput(&createdBean)
}

func (bu *beanUsecase) Update(beanID uint, userID string, data dto.BeanInput, newImageFile *multipart.FileHeader) (dto.BeanOutput, error) {
	// Beanの存在確認と所有者チェック
	var existingBean bean.Entity
	if err := bu.br.GetById(&existingBean, beanID); err != nil {
		return dto.BeanOutput{}, fmt.Errorf("bean not found: %w", err)
	}

	// 所有者チェック
	if existingBean.UserID.String() != userID {
		return dto.BeanOutput{}, fmt.Errorf("access denied: you can only update your own beans")
	}

	// 既存のBeanを更新データで上書き
	bean, varietyIDs := dto_entity.DtoBeanToEntity(userID, data)
	bean.ID = existingBean.ID
	bean.CreatedAt = existingBean.CreatedAt

	// 既存の画像キーを保持（新しい画像がない場合）
	if newImageFile == nil {
		bean.ImageKey = existingBean.ImageKey
	}

	// 新しい画像をS3にアップロード（画像ファイルがある場合）
	if newImageFile != nil {

		//　既存のimageがある場合は削除
		if existingBean.ImageKey != nil {
			if err := bu.s3Service.RemoveBeanImage(*existingBean.ImageKey); err != nil {
				return dto.BeanOutput{}, fmt.Errorf("failed to remove old image: %w", err)
			}
		}

		imageKey, err := bu.s3Service.UploadBeanImage(beanID, newImageFile)
		if err != nil {
			return dto.BeanOutput{}, fmt.Errorf("failed to upload image: %w", err)
		}
		bean.ImageKey = &imageKey
	}

	// Beanを更新
	if err := bu.br.Update(&bean); err != nil {
		return dto.BeanOutput{}, fmt.Errorf("failed to update bean: %w", err)
	}

	// 品種の関連付けを更新
	if err := bu.br.SetVarieties(beanID, varietyIDs); err != nil {
		return dto.BeanOutput{}, fmt.Errorf("failed to update varieties: %w", err)
	}

	// BeanRatingがある場合は作成または更新
	if data.BeanRating != nil {
		if err := bu.handleBeanRating(beanID, userID, data.BeanRating); err != nil {
			return dto.BeanOutput{}, err
		}
	}

	return bu.Read(beanID)
}

func (bu *beanUsecase) validateImageFile(imageFile *multipart.FileHeader) error {
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
func (bu *beanUsecase) validateInputData(userID string, imageFile *multipart.FileHeader) error {
	// ユーザーの存在確認
	var userEntity user.Entity
	if err := bu.ur.GetById(&userEntity, uuid.MustParse(userID)); err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	// 画像ファイルのバリデーション（ファイルがある場合のみ）
	if imageFile != nil {
		if err := bu.validateImageFile(imageFile); err != nil {
			return fmt.Errorf("invalid image file: %w", err)
		}
	}

	return nil
}

// createBeanRating は新しいBeanRatingを作成します
func (bu *beanUsecase) createBeanRating(beanID uint, userID string, ratingData *dto.BeanRatingRef) error {
	beanRatingEntity := bean_rating.Entity{
		BeanID:     beanID,
		UserID:     uuid.MustParse(userID),
		Bitterness: ratingData.Bitterness,
		Acidity:    ratingData.Acidity,
		Body:       ratingData.Body,
	}

	if ratingData.FlavorNote != nil {
		beanRatingEntity.FlavorNote = *ratingData.FlavorNote
	}

	if err := bu.brr.Create(&beanRatingEntity); err != nil {
		return fmt.Errorf("failed to create bean rating: %w", err)
	}

	return nil
}

// handleBeanRating はBeanRatingの作成または更新を処理します
func (bu *beanUsecase) handleBeanRating(beanID uint, userID string, ratingData *dto.BeanRatingRef) error {
	if ratingData.ID != nil {
		// IDがある場合は更新
		beanRatingEntity := bean_rating.Entity{
			ID:         uint(*ratingData.ID),
			BeanID:     beanID,
			UserID:     uuid.MustParse(userID),
			Bitterness: ratingData.Bitterness,
			Acidity:    ratingData.Acidity,
			Body:       ratingData.Body,
		}

		if ratingData.FlavorNote != nil {
			beanRatingEntity.FlavorNote = *ratingData.FlavorNote
		}

		if err := bu.brr.UpdateByID(&beanRatingEntity); err != nil {
			return fmt.Errorf("failed to update bean rating: %w", err)
		}
	} else {
		// IDがない場合は新規作成
		if err := bu.createBeanRating(beanID, userID, ratingData); err != nil {
			return err
		}
	}

	return nil
}

// generateBeanOutput は画像URLを生成してBeanOutputを作成します
func (bu *beanUsecase) generateBeanOutput(domainBean *bean.Entity) (dto.BeanOutput, error) {
	var imageURL string
	if domainBean.ImageKey != nil {
		url, err := bu.s3Service.GenerateBeanImageURL(*domainBean.ImageKey)
		if err != nil {
			return dto.BeanOutput{}, fmt.Errorf("failed to generate image URL: %w", err)
		}
		imageURL = url
	}

	// Convert domain entity directly to DTO
	return dto_entity.EntityBeanToDto(domainBean, imageURL), nil
}

func NewBeanUsecase(ur user.IUserRepository, br bean.IBeanRepository, brr bean_rating.IBeanRatingRepository, s3Service s3.IS3Service) IBeanUsecase {
	return &beanUsecase{
		ur:        ur,
		br:        br,
		brr:       brr,
		s3Service: s3Service,
	}
}
