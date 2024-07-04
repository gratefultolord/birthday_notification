package handlers

import (
	"birthday_notification/internal/repository"
	"birthday_notification/internal/utils"
	"net/http"
)

type BirthdayHandler struct {
	Repo repository.UserRepositoryInterface
}

func NewBirthdayHandler(repo repository.UserRepositoryInterface) *BirthdayHandler {
	return &BirthdayHandler{Repo: repo}
}

func (h *BirthdayHandler) GetBirthdays(w http.ResponseWriter, r *http.Request) {
	users, err := h.Repo.GetAllUsers()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Ошибка получения пользователей")
		return
	}

	var birthdaysToday []string
	for _, user := range users {
		birthdate, err := utils.ParseDate(user.Birthdate)
		if err != nil {
			continue // Игнорировать ошибки парсинга даты
		}
		if utils.IsToday(birthdate) {
			birthdaysToday = append(birthdaysToday, user.Name)
		}
	}

	utils.RespondWithJSON(w, http.StatusOK, birthdaysToday)
}
