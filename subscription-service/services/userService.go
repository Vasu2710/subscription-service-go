package services

import (
	"errors"
	"subscription-service/models"
	"subscription-service/repositories"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func RegisterUser(user models.User) error {
	// Add password hashing logic here if necessary
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err // Return an error if password hashing fails
	}
	user.Password = string(hashedPassword)
	return repositories.CreateUser(user)
}

func AuthenticateUser(email, password string) (models.User, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return models.User{}, err
	}
	// Add password comparison logic here'
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// Passwords don't match
		return models.User{}, errors.New("invalid credentials")
	}
	return user, nil
}
