package controllers

import (
	"mypic/config"
	"mypic/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SignupRequest struct {
	UserName     string `json:"userName"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	UserPassword string `json:"userPassword"`
	UserLogoURL  string `json:"userLogoURL"`
}

type LoginRequest struct {
	Identifier   string `json:"identifier"`
	UserPassword string `json:"userPassword"`
}

type UpdateUserRequest struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	UserLogoURL  string `json:"userLogoURL"`
	UserPassword string `json:"userPassword"`
}

func Signup(c *gin.Context) {
	var req SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := services.Signup(
		req.UserName,
		req.Name,
		req.Email,
		req.UserPassword,
		req.UserLogoURL,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := services.Login(req.Identifier, req.UserPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := config.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
		"user": gin.H{
			"id":       user.ID,
			"userName": user.UserName,
			"name":     user.Name,
			"email":    user.Email,
		},
	})
}

func UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err = services.UpdateUser(
		uint(id),
		req.Name,
		req.Email,
		req.UserLogoURL,
		req.UserPassword,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	err = services.DeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
