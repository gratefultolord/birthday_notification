package server

import (
	"birthday_notification/internal/handlers"
	"birthday_notification/internal/repository"
	"birthday_notification/pkg/middleware"
	"database/sql"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func NewServer(db *sql.DB) *Server {
	r := mux.NewRouter()
	userRepo := repository.NewUserRepository(db)
	subscriptionRepo := repository.NewSubscriptionRepository(db)

	birthdayHandler := handlers.NewBirthdayHandler(userRepo)
	subscriptionHandler := handlers.NewSubscriptionHandler(subscriptionRepo)

	r.HandleFunc("/birthdays", birthdayHandler.GetBirthdays).Methods("GET")
	r.HandleFunc("/subscribe", subscriptionHandler.Subscribe).Methods("POST")
	r.HandleFunc("/unsubscribe", subscriptionHandler.Unsubscribe).Methods("POST")

	r.Use(middleware.AuthMiddleware)

	return &Server{Router: r}
}
