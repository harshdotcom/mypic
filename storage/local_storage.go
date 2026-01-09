package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type LocalStorage struct {
	BasePath string
	BaseURL  string
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		BasePath: "./uploads",
		BaseURL:  "http://localhost:8080/uploads",
	}
}

func (s *LocalStorage) Upload(file *multipart.FileHeader) (storedName, url string, err error) {
	ext := filepath.Ext(file.Filename)
	storedName = uuid.New().String() + ext

	if err := os.MkdirAll(s.BasePath, os.ModePerm); err != nil {
		return "", "", err
	}

	dstPath := filepath.Join(s.BasePath, storedName)

	src, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return "", "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", "", err
	}

	url = fmt.Sprintf("%s/%s", s.BaseURL, storedName)
	return storedName, url, nil
}
