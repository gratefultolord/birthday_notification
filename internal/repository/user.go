package repository

import (
	"birthday_notification/internal/models"
)

type UserRepositoryInterface interface {
	GetAllUsers() ([]models.User, error)
}

type UserRepository struct {
	Repo UserRepositoryInterface
}

// NewUserRepository создает новый экземпляр UserRepository с реальным репозиторием.
func NewUserRepository() *UserRepository {
	return &UserRepository{Repo: &RealUserRepository{}}
}

func (repo *UserRepository) GetAllUsers() ([]models.User, error) {
	return repo.Repo.GetAllUsers()
}

type RealUserRepository struct{}

func (repo *RealUserRepository) GetAllUsers() ([]models.User, error) {
	return []models.User{
		{ID: 1, Name: "Alice", Birthdate: "1990-07-04", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Birthdate: "1988-07-05", Email: "bob@example.com"},
	}, nil
}
