package repositories

import (
	"context"
	"subscription-service/database"
	"subscription-service/models"

	"github.com/google/uuid"
)

func CreateUser(user models.User) error {
	user.ID = uuid.New().String()

	// SQL query to insert user details into the database
	query := `INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)`
	_, err := database.DB.Exec(context.Background(), query, user.ID, user.Name, user.Email, user.Password)
	return err
}

func GetUserByID(userID string) (models.User, error) {
	var user models.User
	query := `SELECT id, name, email, password FROM users WHERE id = $1`
	err := database.DB.QueryRow(context.Background(), query, userID).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return user, err
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	err := database.DB.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return user, err
}
func GetAllUsers() ([]models.User, error) {
	query := `SELECT id, name, email, password FROM users`
	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
