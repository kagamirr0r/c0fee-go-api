package s3

import (
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewS3Client() (*minio.Client, error) {
	var endpoint string
	if os.Getenv("GO_ENV") == "dev" {
		endpoint = "localhost:9000"
	} else {
		endpoint = "s3.amazonaws.com"
	}

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	useSSL := os.Getenv("GO_ENV") != "dev"

	// S3クライアントの初期化
	s3Client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return s3Client, nil
}
