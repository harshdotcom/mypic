package services

import (
	"errors"

	"mypic/models"
	"mypic/repositories"

	"golang.org/x/crypto/bcrypt"
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

	if bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(password)) != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func UpdateUser(id uint, name, email, logo string) error {
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

	return repositories.UpdateUser(user)
}

func DeleteUser(id uint) error {
	user, err := repositories.GetUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	return repositories.DeleteUser(user)
}
