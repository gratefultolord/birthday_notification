package service

import (
	"birthday_notification/internal/models"
	"birthday_notification/internal/repository"
	"birthday_notification/internal/utils"
)

type BirthdayService struct {
	repo *repository.UserRepository
}

func NewBirthdayService(repo *repository.UserRepository) *BirthdayService {
	return &BirthdayService{repo: repo}
}

func (s *BirthdayService) GetBirthdays() ([]models.User, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var birthdaysToday []models.User
	for _, user := range users {
		birthdate, err := utils.ParseDate(user.Birthdate)
		if err != nil {
			continue
		}
		if utils.IsToday(birthdate) {
			birthdaysToday = append(birthdaysToday, user)
		}
	}

	return birthdaysToday, nil
}
