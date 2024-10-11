package repositories

import (
	configs "auraskin/internal/configs/firebase"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/google/uuid"
)

type StorageRepository interface {
	UploadFile(file multipart.File, fileHeader *multipart.FileHeader, folder string) (string, error)
}

type storageRepository struct {
	bucketName string
}

func NewStorageRepository() StorageRepository {
	return &storageRepository{
		bucketName: configs.FirebaseStorageBucketName,
	}
}

func (repo *storageRepository) UploadFile(file multipart.File, fileHeader *multipart.FileHeader, folder string) (string, error) {
	ctx := context.Background()

	client, err := configs.InitializeFirebaseApp().Storage(ctx)
	if err != nil {
		return "", err
	}

	bucket, err := client.Bucket(repo.bucketName)
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("%s/%d_%s", folder, time.Now().Unix(), fileHeader.Filename)

	writer := bucket.Object(fileName).NewWriter(ctx)

	token := generateUUID()
	writer.Metadata = map[string]string{
		"firebaseStorageDownloadTokens": token,
	}
	defer writer.Close()

	if _, err := io.Copy(writer, file); err != nil {
		return "", err
	}

	encodedFileName := url.QueryEscape(fileName)

	fileUrl := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", repo.bucketName, encodedFileName, token)

	return fileUrl, nil
}

func generateUUID() string {
	return uuid.New().String()
}
