package minio

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/vanthang24803/mini/internal/config"
	"github.com/vanthang24803/mini/pkg/logger"
)

var minioClient *minio.Client

func Init() {
	cfg := config.New()
	log := logger.GetLogger()
	ctx := context.Background()

	endpoint := cfg.Minio.Endpoint
	accessKeyID := cfg.Minio.AccessKeyID
	secretAccessKey := cfg.Minio.SecretAccessKey
	useSSL := cfg.Minio.UseSSL
	bucketName := cfg.Minio.BucketName

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Error(fmt.Sprintf("error for connect minio storage %v", err))
	}

	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to check if bucket exists: %v", err))
	}

	if !exists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatal(fmt.Sprintf("failed to create bucket: %v", err))
		}
		log.Info(fmt.Sprintf("Created new bucket: %s", bucketName))
	} else {
		log.Info(fmt.Sprintf("Connected to existing bucket: %s", bucketName))
	}
}

func UploadFile(objectName string, file []byte) error {
	ctx := context.Background()
	bucketName := config.New().Minio.BucketName
	_, err := minioClient.PutObject(ctx, bucketName, objectName, bytes.NewReader(file), int64(len(file)), minio.PutObjectOptions{})

	return err
}

func RemoveFile(objectName string) error {
	ctx := context.Background()
	bucketName := config.New().Minio.BucketName
	err := minioClient.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

func DownloadFile(objectName string) ([]byte, error) {
	ctx := context.Background()
	bucketName := config.New().Minio.BucketName
	object, err := minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()

	file, err := io.ReadAll(object)
	if err != nil {
		return nil, err
	}

	return file, nil
}
