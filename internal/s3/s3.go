package s3

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3Client struct {
	Client *minio.Client
	Bucket string
}

func NewS3Client(endpoint, accessKeyID, secretAccessKey, bucket string) (*S3Client, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		return nil, err
	}

	return &S3Client{
		Client: client,
		Bucket: bucket,
	}, nil
}

func (s *S3Client) PutObject(src io.Reader, fileSize int64, fileName string, contentType string) (minio.UploadInfo, error) {
	info, err := s.Client.PutObject(context.Background(), s.Bucket, fileName, src, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return minio.UploadInfo{}, err
	}

	return info, nil
}

func (s *S3Client) GetObject(fileName string) (*minio.Object, error) {
	obj, err := s.Client.GetObject(context.Background(), s.Bucket, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	return obj, nil
}
