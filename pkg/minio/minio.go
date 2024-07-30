package minio

import (
	"context"
	"errors"
	"time"
)

func GetFileTemporaryURL(bucketName, objectName string) (string, error) {
	if len(bucketName) <= 0 || len(objectName) <= 0 {
		return "", errors.New("invalid argument")
	}

	expiry := time.Second * time.Duration(ExpireTime)

	presignedURL, err := MinioClient.PresignedGetObject(context.Background(), bucketName, objectName, expiry, nil)

	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}