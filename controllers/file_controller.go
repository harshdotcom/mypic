package controllers

import (
	"net/http"
	"strconv"

	"mypic/services"

	"github.com/gin-gonic/gin"
)

type ListFilesRequest struct {
	Search string `json:"search"`
	SortBy string `json:"sortBy"` // "name" or "date"
	Order  string `json:"order"`  // "asc" or "desc"
}

// Upload multiple files
func UploadFiles(c *gin.Context) {
	userID := c.GetUint("userId")

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid form"})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no files provided"})
		return
	}

	result, err := services.UploadFiles(userID, files)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// List user files (POST-based)
func ListFiles(c *gin.Context) {
	userID := c.GetUint("userId")

	var req ListFilesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	files, err := services.ListUserFiles(userID, req.Search, req.SortBy, req.Order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, files)
}

func DeleteFile(c *gin.Context) {
	userID := c.GetUint("userId")

	idParam := c.Param("id")
	fileID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file id"})
		return
	}

	err = services.DeleteUserFile(userID, uint(fileID))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "file deleted successfully"})
}
