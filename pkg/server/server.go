package server

import (
	"birthday_notification/internal/handlers"
	"birthday_notification/internal/repository"
	"birthday_notification/pkg/middleware"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func NewServer() *Server {
	r := mux.NewRouter()
	repo := repository.NewUserRepository()
	birthdayHandler := handlers.NewBirthdayHandler(repo.Repo)

	r.HandleFunc("/birthdays", birthdayHandler.GetBirthdays).Methods("GET")
	r.HandleFunc("/subscribe", handlers.Subscribe).Methods("POST")
	r.HandleFunc("/unsubscribe", handlers.Unsubscribe).Methods("POST")

	// Применение middleware для всех маршрутов
	r.Use(middleware.AuthMiddleware)

	return &Server{Router: r}
}
