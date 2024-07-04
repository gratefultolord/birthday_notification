package auth

import (
	"net/http"
)

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Реализация авторизации
		next.ServeHTTP(w, r)
	})
}
