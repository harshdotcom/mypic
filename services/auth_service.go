package services

import (
	"context"
	"errors"
	"os"

	"mypic/models"
	"mypic/repositories"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"
)

func Signup(userName, name, email, password, logo string) error {
	if userName == "" || name == "" || email == "" || password == "" {
		return errors.New("username, name, email and password are required")
	}

	if _, err := repositories.GetUserByUserName(userName); err == nil {
		return errors.New("username already exists")
	}

	if _, err := repositories.GetUserByEmail(email); err == nil {
		return errors.New("email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		UserName:     userName,
		Name:         name,
		Email:        email,
		UserPassword: string(hash),
		UserLogoURL:  logo,
		AuthProvider: "LOCAL", // ✅ THIS IS THE IMPORTANT ADDITION
	}

	return repositories.CreateUser(&user)
}

func Login(identifier, password string) (*models.User, error) {
	if identifier == "" || password == "" {
		return nil, errors.New("email/username and password are required")
	}

	var user *models.User
	var err error

	// Try email first
	user, err = repositories.GetUserByEmail(identifier)
	if err != nil {
		// Try username if email not found
		user, err = repositories.GetUserByUserName(identifier)
		if err != nil {
			return nil, errors.New("invalid credentials")
		}
	}

	// ✅ NEW CHECK (important)
	if user.AuthProvider == "GOOGLE" {
		return nil, errors.New("please login using Google")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(password)) != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func UpdateUser(id uint, name, email, logo string, password string) error {
	user, err := repositories.GetUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	if logo != "" {
		user.UserLogoURL = logo
	}
	if password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.UserPassword = string(hash)
	}

	return repositories.UpdateUser(user)
}

func DeleteUser(id uint) error {
	user, err := repositories.GetUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	return repositories.DeleteUser(user)
}

func GoogleLogin(idToken string) (*models.User, error) {
	if idToken == "" {
		return nil, errors.New("id token is required")
	}

	// Verify token with Google
	payload, err := idtoken.Validate(context.Background(), idToken, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		return nil, errors.New("invalid google token")
	}

	// Extract user info from token
	email, ok := payload.Claims["email"].(string)
	if !ok || email == "" {
		return nil, errors.New("email not found in google token")
	}

	name, _ := payload.Claims["name"].(string)
	picture, _ := payload.Claims["picture"].(string)

	// Check if user already exists
	user, err := repositories.GetUserByEmail(email)
	if err == nil {
		// User exists → return it
		return user, nil
	}

	// User does not exist → create new one
	newUser := models.User{
		UserName:     email, // or generate from email
		Name:         name,
		Email:        email,
		UserLogoURL:  picture,
		UserPassword: "",       // no password for Google users
		AuthProvider: "GOOGLE", // IMPORTANT
	}

	if err := repositories.CreateUser(&newUser); err != nil {
		return nil, err
	}

	return &newUser, nil
}
