package repositories

import (
	"mypic/config"
	"mypic/models"
)

func SaveFile(file *models.File) error {
	return config.DB.Create(file).Error
}

func ListFilesByUser(
	userID uint,
	search string,
	extension string,
	mimeType string,
	startDate string,
	endDate string,
	sortBy string,
	order string,
) ([]models.File, error) {

	var files []models.File
	query := config.DB.Where("user_id = ?", userID)

	if search != "" {
		query = query.Where("original_name LIKE ?", "%"+search+"%")
	}

	if extension != "" {
		query = query.Where("extension = ?", extension)
	}

	if mimeType != "" {
		query = query.Where("mime_type = ?", mimeType)
	}

	if startDate != "" && endDate != "" {
		query = query.Where(
			"created_at BETWEEN ? AND ?",
			startDate+" 00:00:00",
			endDate+" 23:59:59",
		)
	} else if startDate != "" {
		query = query.Where("created_at >= ?", startDate+" 00:00:00")
	} else if endDate != "" {
		query = query.Where("created_at <= ?", endDate+" 23:59:59")
	}

	// Whitelist allowed sort fields
	sortColumn := "created_at" // default

	switch sortBy {
	case "name":
		sortColumn = "original_name"
	case "date":
		sortColumn = "created_at"
	case "size":
		sortColumn = "size"
	case "type":
		sortColumn = "extension"
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
