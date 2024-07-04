package handlers

import (
	"birthday_notification/internal/repository"
	"birthday_notification/internal/utils"
	"log"
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

	log.Printf("Полученные пользователи: %+v\n", users)

	var birthdaysToday []string
	for _, user := range users {
		birthdate, err := utils.ParseDate(user.Birthdate)
		if err != nil {
			log.Printf("Ошибка парсинга даты для пользователя %s: %v\n", user.Name, err)
			continue
		}
		if utils.IsToday(birthdate) {
			birthdaysToday = append(birthdaysToday, user.Name)
		}
	}

	log.Printf("Дни рождения сегодня: %+v\n", birthdaysToday)

	utils.RespondWithJSON(w, http.StatusOK, birthdaysToday)
}
