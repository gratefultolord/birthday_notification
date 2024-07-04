package handlers

import (
	"net/http"
)

func Subscribe(w http.ResponseWriter, r *http.Request) {
	// Реализация подписки
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Подписка успешна"))
}

func Unsubscribe(w http.ResponseWriter, r *http.Request) {
	// Реализация отписки
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Отписка успешна"))
}
