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
	"strings"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type IBeanUsecase interface {
	Read(bean model.Bean) (dto.BeanOutput, error)
	Create(userID string, dataJSON string, imageFile *multipart.FileHeader) (dto.BeanOutput, error)
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

	return converter.ConvertToBeanResponse(&storedBean, imageURL), nil
}

func (bu *beanUsecase) Create(userID string, dataJSON string, imageFile *multipart.FileHeader) (dto.BeanOutput, error) {
	// ユーザーの存在確認
	var user model.User
	if err := bu.ur.GetById(&user, uuid.MustParse(userID)); err != nil {
		return dto.BeanOutput{}, fmt.Errorf("user not found: %w", err)
	}

	// JSONデータをパース
	var data dto.CreateBeanData
	if err := json.Unmarshal([]byte(dataJSON), &data); err != nil {
		return dto.BeanOutput{}, fmt.Errorf("invalid JSON data: %w", err)
	}
	// validatorを使用してJsonデータをバリデーション
	if err := bu.validator.Struct(data); err != nil {
		return dto.BeanOutput{}, fmt.Errorf("validation failed: %w", err)
	}

	// 画像ファイルのバリデーション（ファイルがある場合のみ）
	if imageFile != nil {
		if err := bu.validateImageFile(imageFile); err != nil {
			return dto.BeanOutput{}, err
		}
	}

	// Beanエンティティを作成
	bean := converter.ConvertCreateBeanDataToBean(userID, data)
	// 最初にBeanを保存（画像なしで）
	if err := bu.br.Create(&bean); err != nil {
		return dto.BeanOutput{}, fmt.Errorf("failed to create bean: %w", err)
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
		beanRating := model.BeanRating{
			BeanID:     bean.ID,
			UserID:     uuid.MustParse(userID),
			Bitterness: data.BeanRating.Bitterness,
			Acidity:    data.BeanRating.Acidity,
			Body:       data.BeanRating.Body,
		}

		if data.BeanRating.FlavorNote != nil {
			beanRating.FlavorNote = *data.BeanRating.FlavorNote
		}

		if err := bu.brr.Create(&beanRating); err != nil {
			return dto.BeanOutput{}, fmt.Errorf("failed to create bean rating: %w", err)
		}

		// BeanRatingを含めて再取得
		if err := bu.br.GetById(&createdBean, bean.ID); err != nil {
			return dto.BeanOutput{}, fmt.Errorf("failed to get created bean with rating: %w", err)
		}
	}

	// 画像URLを生成
	var imageURL string
	if createdBean.ImageKey != nil {
		imageURL, _ = bu.s3Service.GenerateBeanImageURL(*createdBean.ImageKey)
	}

	return converter.ConvertToBeanResponse(&createdBean, imageURL), nil
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

	valid := false
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			valid = true
			break
		}
	}

	if !valid {
		return errors.New("image file must be jpg, jpeg, png, or webp")
	}

	return nil
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
