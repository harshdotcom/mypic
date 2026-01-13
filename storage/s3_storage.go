package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type S3Storage struct {
	Client *s3.Client
	Bucket string
	Region string
}

func NewS3Storage() (*S3Storage, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		return nil, err
	}

	return &S3Storage{
		Client: s3.NewFromConfig(cfg),
		Bucket: os.Getenv("AWS_S3_BUCKET"),
		Region: os.Getenv("AWS_REGION"),
	}, nil
}

func (s *S3Storage) Upload(file *multipart.FileHeader) (storedName, url string, err error) {
	ext := filepath.Ext(file.Filename)
	storedName = uuid.New().String() + ext

	src, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer src.Close()

	_, err = s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(storedName),
		Body:   src,
	})

	if err != nil {
		return "", "", err
	}

	url = fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.Bucket, s.Region, storedName)
	return storedName, url, nil
}

func (s *S3Storage) Delete(storedName string) error {
	_, err := s.Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(storedName),
	})
	return err
}
