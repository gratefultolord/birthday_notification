package main

import (
	"birthday_notification/config"
	"birthday_notification/pkg/server"
	"log"
	"net/http"
)

func main() {
	// Инициализация соединения с базой данных
	config.InitDB()
	defer config.DB.Close()

	srv := server.NewServer(config.DB)
	log.Fatal(http.ListenAndServe(":8080", srv.Router))
}
