// infrastructure/s3/service.go
package s3

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

// TODO: 以下はアプリケーションサービスのInterfaceとして切り出す
type IS3Service interface {
	GeneratePresignedURL(bucket, objectKey string, expiry time.Duration) (string, error)
	GenerateBeanImageURL(imageKey string) (string, error)
	GenerateUserAvatarURL(avatarKey string) (string, error)
	GenerateRoasterImageURL(imageKey string) (string, error)
	UploadBeanImage(beanID uint, imageFile *multipart.FileHeader) (string, error)
}

type s3Service struct {
	client *minio.Client
}

func (s *s3Service) GeneratePresignedURL(bucket, objectKey string, expiry time.Duration) (string, error) {
	if objectKey == "" || objectKey == "null" {
		return "", nil
	}

	presignedURL, err := s.client.PresignedGetObject(
		context.Background(),
		bucket,
		objectKey,
		expiry,
		nil,
	)

	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

func (s *s3Service) GenerateBeanImageURL(imageKey string) (string, error) {
	return s.generateImageURL("beans", imageKey)
}

func (s *s3Service) GenerateUserAvatarURL(avatarKey string) (string, error) {
	return s.generateImageURL("users", avatarKey)
}

func (s *s3Service) GenerateRoasterImageURL(imageKey string) (string, error) {
	return s.generateImageURL("roasters", imageKey)
}

// 共通の画像URL生成関数
func (s *s3Service) generateImageURL(prefix, imageKey string) (string, error) {
	if imageKey == "" || imageKey == "null" {
		return "", nil
	}

	bucket := os.Getenv("S3_BUCKET")
	objectKey := prefix + "/" + imageKey

	return s.GeneratePresignedURL(bucket, objectKey, time.Hour*1)
}

func (s *s3Service) UploadBeanImage(beanID uint, imageFile *multipart.FileHeader) (string, error) {
	if imageFile == nil {
		return "", fmt.Errorf("image file is required")
	}

	bucket := os.Getenv("S3_BUCKET")
	if bucket == "" {
		return "", fmt.Errorf("S3_BUCKET environment variable is not set")
	}

	// ファイルを開く
	file, err := imageFile.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open image file: %w", err)
	}
	defer file.Close()

	// ファイル拡張子を取得
	ext := strings.ToLower(filepath.Ext(imageFile.Filename))

	// ユニークなファイル名を生成
	uniqueID := uuid.New().String()
	imageKey := fmt.Sprintf("%d/%s_%s", beanID, uniqueID, ext)
	fmt.Println("Generated image key:", imageKey)
	objectKey := "beans/" + imageKey

	// Content-Typeを設定
	var contentType string
	switch ext {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".webp":
		contentType = "image/webp"
	default:
		contentType = "application/octet-stream"
	}

	// S3互換ストレージにアップロード（MinIO/S3自動判定）
	_, err = s.client.PutObject(
		context.Background(),
		bucket,
		objectKey,
		file,
		imageFile.Size,
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to upload image: %w", err)
	}

	return imageKey, nil
}

func NewS3Service() (IS3Service, error) {
	client, err := NewS3Client()
	if err != nil {
		return nil, err
	}
	return &s3Service{client: client}, nil
}
