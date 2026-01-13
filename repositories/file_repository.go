package repositories

import (
	"mypic/config"
	"mypic/models"
)

// Save uploaded file metadata
func SaveFile(file *models.File) error {
	return config.DB.Create(file).Error
}

// List files for a user with optional search and sorting
func ListFilesByUser(
	userID uint,
	search string,
	sortBy string,
	order string,
) ([]models.File, error) {

	var files []models.File
	query := config.DB.Where("user_id = ?", userID)

	if search != "" {
		query = query.Where("original_name LIKE ?", "%"+search+"%")
	}

	// Whitelist allowed sort fields
	sortColumn := "created_at"
	if sortBy == "name" {
		sortColumn = "original_name"
	}

	// Default sort order
	sortOrder := "desc"
	if order == "asc" {
		sortOrder = "asc"
	}

	query = query.Order(sortColumn + " " + sortOrder)

	err := query.Find(&files).Error
	return files, err
}

func GetFileByID(id uint) (*models.File, error) {
	var file models.File
	err := config.DB.First(&file, id).Error
	return &file, err
}

func DeleteFile(file *models.File) error {
	return config.DB.Delete(file).Error
}
