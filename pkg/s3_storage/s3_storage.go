package s3storage

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// go get github.com/aws/aws-sdk-go-v2
// go get github.com/aws/aws-sdk-go-v2/config
// go get github.com/aws/aws-sdk-go-v2/service/s3
// go get github.com/aws/aws-sdk-go-v2/feature/s3/manager bunlar yüklendi

type S3Config struct {
	BucketName string
	Key        string
	Body       []byte
	ACL        types.ObjectCannedACL
}

type S3Service struct {
	Client *s3.Client
}

func NewS3Service() (*S3Service, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)
	return &S3Service{Client: client}, nil
}

func (s *S3Service) UploadFile(cfg S3Config) (string, error) {
	uploader := manager.NewUploader(s.Client)

	_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(cfg.BucketName),
		Key:    aws.String(cfg.Key),
		Body:   bytes.NewReader(cfg.Body),
		ACL:    cfg.ACL,
	})
	if err != nil {
		return "", err
	}

	fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", cfg.BucketName, cfg.Key) //burdaki urli dbye kaydetcem
	return fileURL, nil
}
