package fileuploadersvc

import (
	"context"
	"fmt"

	"github.com/arfan21/project-sprint-banking-api/config"
	"github.com/arfan21/project-sprint-banking-api/internal/model"
	"github.com/arfan21/project-sprint-banking-api/pkg/constant"
	"github.com/arfan21/project-sprint-banking-api/pkg/s3"
	"github.com/arfan21/project-sprint-banking-api/pkg/validation"
)

type Service struct {
	s3Client *s3.S3
}

func New() *Service {
	client, _ := s3.New()
	return &Service{
		s3Client: client,
	}
}

func (s *Service) UploadImage(ctx context.Context, req model.FileUploaderImageRequest) (res model.FileUploaderImageResponse, err error) {
	if req.File == nil {
		err = fmt.Errorf("fileuploader.service.Upload: file is required, %w", constant.ErrFileRequired)
		return
	}
	fieldName := "file"

	err = validation.ValidateContentType(fieldName, req.File.Header.Get("Content-Type"), validation.WithValidateContentTypeImage())
	if err != nil {
		err = fmt.Errorf("fileuploader.service.Upload: failed to validate content type: %w", err)
		return
	}
	// 10KB
	minSize := 10 * 1024
	err = validation.ValidateFileSize(fieldName, req.File.Size, validation.WithValidateFileSizeMinSize(int64(minSize)))
	if err != nil {
		err = fmt.Errorf("imageuploader.service.Upload: failed to validate file size: %w", err)
		return
	}

	folder := "images"
	bucket := config.Get().S3.Bucket
	resStr, err := s.s3Client.Upload(ctx, bucket, folder, req.File)
	if err != nil {
		err = fmt.Errorf("imageuploader.service.Upload: failed to upload file: %w", err)
		return
	}
	url := s.s3Client.GetURL(bucket, resStr)

	res.ImageURL = url

	return res, nil
}
