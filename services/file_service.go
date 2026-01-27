package services

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"

	"mypic/models"
	"mypic/repositories"
	"mypic/storage"
)

var allowedExtensions = []string{".png", ".jpg", ".jpeg", ".pdf", ".txt"}

func UploadFiles(userID uint, files []*multipart.FileHeader) ([]models.File, error) {
	// storage := storage.NewLocalStorage()
	storage, err := storage.NewS3Storage()
	if err != nil {
		return nil, err
	}
	var result []models.File

	for _, file := range files {
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if !isAllowed(ext) {
			continue
		}

		storedName, url, err := storage.Upload(file)
		if err != nil {
			return nil, err
		}

		f := models.File{
			UserID:       userID,
			OriginalName: file.Filename,
			StoredName:   storedName,
			Extension:    ext,
			MimeType:     file.Header.Get("Content-Type"),
			Size:         file.Size,
			URL:          url,
		}

		repositories.SaveFile(&f)
		result = append(result, f)
	}

	return result, nil
}

func isAllowed(ext string) bool {
	for _, a := range allowedExtensions {
		if a == ext {
			return true
		}
	}
	return false
}

func ListUserFiles(
	userID uint,
	search string,
	extension string,
	mimeType string,
	startDate string,
	endDate string,
	sortBy string,
	order string,
) ([]models.File, error) {

	return repositories.ListFilesByUser(
		userID,
		search,
		extension,
		mimeType,
		startDate,
		endDate,
		sortBy,
		order,
	)
}

func DeleteUserFile(userID uint, fileID uint) error {
	file, err := repositories.GetFileByID(fileID)
	if err != nil {
		return err
	}

	if file.UserID != userID {
		return errors.New("unauthorized")
	}

	storage, err := storage.NewS3Storage()
	if err != nil {
		return err
	}

	if err := storage.Delete(file.StoredName); err != nil {
		return err
	}

	return repositories.DeleteFile(file)
}
