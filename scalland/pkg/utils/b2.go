package utils

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// B2Client wraps an S3 client for Backblaze B2.
type B2Client struct {
	client *s3.Client
	bucket string
}

// NewB2Client creates a new B2/S3 client from config.
func NewB2Client(cfg ObjectStorageConfig) (*B2Client, error) {
	resolver := aws.EndpointResolverWithOptionsFunc(
		func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{URL: cfg.Endpoint}, nil
		},
	)
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretKey, "")),
		config.WithEndpointResolverWithOptions(resolver),
	)
	if err != nil {
		return nil, fmt.Errorf("load B2 config: %w", err)
	}
	return &B2Client{client: s3.NewFromConfig(awsCfg), bucket: cfg.Bucket}, nil
}

// Upload uploads a file to B2.
func (b *B2Client) Upload(key string, reader io.Reader, contentType string) error {
	_, err := b.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(b.bucket),
		Key:         aws.String(key),
		Body:        reader,
		ContentType: aws.String(contentType),
	})
	return err
}

// UploadFile uploads a local file to B2.
func (b *B2Client) UploadFile(key, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	return b.Upload(key, f, "application/octet-stream")
}

// PresignedURL generates a pre-signed download URL.
func (b *B2Client) PresignedURL(key string, expiry time.Duration) (string, error) {
	presigner := s3.NewPresignClient(b.client)
	req, err := presigner.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(b.bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(expiry))
	if err != nil {
		return "", err
	}
	return req.URL, nil
}
