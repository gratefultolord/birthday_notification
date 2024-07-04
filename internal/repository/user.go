package repository

import (
	"birthday_notification/internal/models"
	"database/sql"
)

type UserRepositoryInterface interface {
	GetAllUsers() ([]models.User, error)
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := repo.DB.Query("SELECT id, name, birthdate, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Birthdate, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
