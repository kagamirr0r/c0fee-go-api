package usecase

import (
	"c0fee-api/common/converter"
	"c0fee-api/dto"
	"c0fee-api/infrastructure/s3"
	"c0fee-api/model"
	"c0fee-api/repository"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"slices"
	"strings"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type IBeanUsecase interface {
	Read(bean model.Bean) (dto.BeanOutput, error)
	Create(userID string, dataJSON string, imageFile *multipart.FileHeader) (dto.BeanOutput, error)
	Update(beanID uint, userID string, dataJSON string, imageFile *multipart.FileHeader) (dto.BeanOutput, error)
}

type beanUsecase struct {
	ur        repository.IUserRepository
	br        repository.IBeanRepository
	brr       repository.IBeanRatingRepository
	s3Service s3.IS3Service
	validator *validator.Validate
}

func (bu *beanUsecase) Read(bean model.Bean) (dto.BeanOutput, error) {
	storedBean := model.Bean{}
	if err := bu.br.GetById(&storedBean, bean.ID); err != nil {
		return dto.BeanOutput{}, err
	}

	var imageURL string
	if storedBean.ImageKey != nil {
		url, err := bu.s3Service.GenerateBeanImageURL(*storedBean.ImageKey)
		if err != nil {
			return dto.BeanOutput{}, err
		}
		imageURL = url
	}

	return converter.ConvertBeanToOutput(&storedBean, imageURL), nil
}

func (bu *beanUsecase) Create(userID string, dataJSON string, imageFile *multipart.FileHeader) (dto.BeanOutput, error) {
	// 共通バリデーション
	data, err := bu.validateInputData(userID, dataJSON, imageFile)
	if err != nil {
		return dto.BeanOutput{}, err
	}

	// Beanエンティティを作成
	bean, varietyIDs := converter.ConvertBeanInputToBean(userID, data)
	// 最初にBeanを保存（画像なしで）
	if err := bu.br.Create(&bean); err != nil {
		return dto.BeanOutput{}, fmt.Errorf("failed to create bean: %w", err)
	}

	// 品種の関連付け
	if len(varietyIDs) > 0 {
		if err := bu.br.SetVarieties(bean.ID, varietyIDs); err != nil {
			return dto.BeanOutput{}, fmt.Errorf("failed to set varieties: %w", err)
		}
	}

	// 画像をS3にアップロード（画像ファイルがある場合のみ）
	if imageFile != nil {
		imageKey, err := bu.s3Service.UploadBeanImage(bean.ID, imageFile)
		if err != nil {
			return dto.BeanOutput{}, fmt.Errorf("failed to upload image: %w", err)
		}
		bean.ImageKey = &imageKey

		// 画像キーを更新
		if err := bu.br.Update(&bean); err != nil {
			return dto.BeanOutput{}, fmt.Errorf("failed to update bean with image key: %w", err)
		}
	}

	// 作成されたBeanを取得（関連データ含む）
	var createdBean model.Bean
	if err := bu.br.GetById(&createdBean, bean.ID); err != nil {
		return dto.BeanOutput{}, fmt.Errorf("failed to get created bean: %w", err)
	}

	// BeanRatingがある場合は作成
	if data.BeanRating != nil {
		if err := bu.createBeanRating(bean.ID, userID, data.BeanRating); err != nil {
			return dto.BeanOutput{}, err
		}

		// BeanRatingを含めて再取得
		if err := bu.br.GetById(&createdBean, bean.ID); err != nil {
			return dto.BeanOutput{}, fmt.Errorf("failed to get created bean with rating: %w", err)
		}
	}

	return bu.generateBeanOutput(&createdBean)
}

func (bu *beanUsecase) Update(beanID uint, userID string, dataJSON string, imageFile *multipart.FileHeader) (dto.BeanOutput, error) {
	// 共通バリデーション
	data, err := bu.validateInputData(userID, dataJSON, imageFile)
	if err != nil {
		return dto.BeanOutput{}, err
	}

	// Beanの存在確認と所有者チェック
	var existingBean model.Bean
	if err := bu.br.GetById(&existingBean, beanID); err != nil {
		return dto.BeanOutput{}, fmt.Errorf("bean not found: %w", err)
	}

	// 所有者チェック
	if existingBean.UserID.String() != userID {
		return dto.BeanOutput{}, fmt.Errorf("access denied: you can only update your own beans")
	}
	// 既存のBeanを更新データで上書き
	beanToUpdate, varietyIDs := converter.ConvertBeanInputToBean(userID, data) // IDは変更しない
	beanToUpdate.CreatedAt = existingBean.CreatedAt                            // 作成日時は保持

	// 既存の画像キーを保持（新しい画像がない場合）
	if imageFile == nil {
		beanToUpdate.ImageKey = existingBean.ImageKey
	}

	// 新しい画像をS3にアップロード（画像ファイルがある場合）
	if imageFile != nil {
		// 古い画像を削除（存在する場合）
		if existingBean.ImageKey != nil {
			// TODO: S3から古い画像を削除するメソッドを実装する場合は、ここで呼び出す
		}

		imageKey, err := bu.s3Service.UploadBeanImage(beanID, imageFile)
		if err != nil {
			return dto.BeanOutput{}, fmt.Errorf("failed to upload image: %w", err)
		}
		beanToUpdate.ImageKey = &imageKey
	}

	// Beanを更新
	if err := bu.br.Update(&beanToUpdate); err != nil {
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

	// 更新されたBeanを取得（関連データ含む）
	var updatedBean model.Bean
	if err := bu.br.GetById(&updatedBean, beanID); err != nil {
		return dto.BeanOutput{}, fmt.Errorf("failed to get updated bean: %w", err)
	}

	return bu.generateBeanOutput(&updatedBean)
}

func (bu *beanUsecase) validateImageFile(imageFile *multipart.FileHeader) error {
	// ファイルサイズチェック（例：5MB制限）
	maxSize := int64(5 * 1024 * 1024) // 5MB
	if imageFile.Size > maxSize {
		return errors.New("image file size must be less than 5MB")
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
func (bu *beanUsecase) validateInputData(userID string, dataJSON string, imageFile *multipart.FileHeader) (dto.BeanInput, error) {
	// ユーザーの存在確認
	var user model.User
	if err := bu.ur.GetById(&user, uuid.MustParse(userID)); err != nil {
		return dto.BeanInput{}, fmt.Errorf("user not found: %w", err)
	}

	// JSONデータをパース
	var data dto.BeanInput
	if err := json.Unmarshal([]byte(dataJSON), &data); err != nil {
		return dto.BeanInput{}, fmt.Errorf("invalid JSON data: %w", err)
	}

	// validatorを使用してJsonデータをバリデーション
	if err := bu.validator.Struct(data); err != nil {
		return dto.BeanInput{}, fmt.Errorf("validation failed: %w", err)
	}

	// 画像ファイルのバリデーション（ファイルがある場合のみ）
	if imageFile != nil {
		if err := bu.validateImageFile(imageFile); err != nil {
			return dto.BeanInput{}, err
		}
	}

	return data, nil
}

// createBeanRating は新しいBeanRatingを作成します
func (bu *beanUsecase) createBeanRating(beanID uint, userID string, ratingData *dto.BeanRatingRef) error {
	beanRating := model.BeanRating{
		BeanID:     beanID,
		UserID:     uuid.MustParse(userID),
		Bitterness: ratingData.Bitterness,
		Acidity:    ratingData.Acidity,
		Body:       ratingData.Body,
	}

	if ratingData.FlavorNote != nil {
		beanRating.FlavorNote = *ratingData.FlavorNote
	}

	if err := bu.brr.Create(&beanRating); err != nil {
		return fmt.Errorf("failed to create bean rating: %w", err)
	}

	return nil
}

// handleBeanRating はBeanRatingの作成または更新を処理します
func (bu *beanUsecase) handleBeanRating(beanID uint, userID string, ratingData *dto.BeanRatingRef) error {
	if ratingData.ID != nil {
		// IDがある場合は更新
		beanRating := model.BeanRating{
			ID:         uint(*ratingData.ID),
			BeanID:     beanID,
			UserID:     uuid.MustParse(userID),
			Bitterness: ratingData.Bitterness,
			Acidity:    ratingData.Acidity,
			Body:       ratingData.Body,
		}

		if ratingData.FlavorNote != nil {
			beanRating.FlavorNote = *ratingData.FlavorNote
		}

		if err := bu.brr.UpdateByID(&beanRating); err != nil {
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
func (bu *beanUsecase) generateBeanOutput(bean *model.Bean) (dto.BeanOutput, error) {
	var imageURL string
	if bean.ImageKey != nil {
		url, err := bu.s3Service.GenerateBeanImageURL(*bean.ImageKey)
		if err != nil {
			return dto.BeanOutput{}, fmt.Errorf("failed to generate image URL: %w", err)
		}
		imageURL = url
	}

	return converter.ConvertBeanToOutput(bean, imageURL), nil
}

func NewBeanUsecase(ur repository.IUserRepository, br repository.IBeanRepository, brr repository.IBeanRatingRepository, s3Service s3.IS3Service, v *validator.Validate) IBeanUsecase {
	return &beanUsecase{
		ur:        ur,
		br:        br,
		brr:       brr,
		s3Service: s3Service,
		validator: v,
	}
}
