package file_storage

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/romanzimoglyad/inquiry-backend/internal/logger"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/romanzimoglyad/inquiry-backend/internal/config"
)

type S3 struct {
	sess       *session.Session
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
	svc        *s3.S3
}

func NewS3() (*S3, error) {

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewEnvCredentials(),
		Region:      aws.String(config.Config.S3.Region)},
	)

	if err != nil {
		return nil, fmt.Errorf("error in NewSession: %w", err)
	}

	return &S3{svc: s3.New(sess), sess: sess, uploader: s3manager.NewUploader(sess), downloader: s3manager.NewDownloader(sess)}, nil
}

func (s *S3) Upload(file *domain.File) error {

	_, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.Config.S3.Bucket),
		Key:    aws.String(file.Name),
		Body:   bytes.NewReader(file.Data),
	})
	if err != nil {
		// Print the error and exit.
		return fmt.Errorf("error in NewUploader: %w", err)
	}

	logger.Info().Msgf("Successfully uploaded %q\n", file.Name)
	return nil
}

func (s *S3) GetUrl(key string) (string, error) {
	req, _ := s.svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(config.Config.S3.Bucket),
		Key:    aws.String(key),
	})
	url, err := req.Presign(167 * time.Hour) // Set the expiration time for the URL (15 minutes in this example)
	if err != nil {

		return "", fmt.Errorf("error generating pre-signed URL:", err)
	}

	return url, nil
}

func (s *S3) Download(key string) (*domain.File, error) {

	result, err := s.svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(config.Config.S3.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("error in Download: %w", err)
	}
	defer result.Body.Close()

	buffer := bytes.Buffer{}
	_, err = buffer.ReadFrom(result.Body)
	if err != nil {

		return nil, fmt.Errorf("Error reading file content:", err)
	}

	return &domain.File{
		Name: key,
		Data: buffer.Bytes(),
	}, nil
}
