// infrastructure/s3/service.go
package s3

import (
	"context"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
)

type IS3Service interface {
	GeneratePresignedURL(bucket, objectKey string, expiry time.Duration) (string, error)
	GenerateBeanImageURL(imageKey string) (string, error)
	GenerateUserAvatarURL(avatarKey string) (string, error)
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
	if imageKey == "" || imageKey == "null" {
		return "", nil
	}

	bucket := os.Getenv("S3_BUCKET")
	objectKey := "beans/" + imageKey

	return s.GeneratePresignedURL(bucket, objectKey, time.Hour*1)
}

func (s *s3Service) GenerateUserAvatarURL(avatarKey string) (string, error) {
	if avatarKey == "" || avatarKey == "null" {
		return "", nil
	}

	bucket := os.Getenv("S3_BUCKET")
	objectKey := "users/" + avatarKey

	return s.GeneratePresignedURL(bucket, objectKey, time.Hour*1)
}

func NewS3Service() (IS3Service, error) {
	client, err := NewS3Client()
	if err != nil {
		return nil, err
	}
	return &s3Service{client: client}, nil
}
