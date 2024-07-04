package main

import (
	"birthday_notification/pkg/server"
	"log"
	"net/http"
)

func main() {
	srv := server.NewServer()
	log.Fatal(http.ListenAndServe(":8080", srv.Router))
}
