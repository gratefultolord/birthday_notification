package tests

import (
	"birthday_notification/pkg/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	rr := httptest.NewRecorder()

	authMiddleware := middleware.AuthMiddleware(handler)

	// Test without Authorization header
	authMiddleware.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusForbidden)
	}

	// Test with invalid Authorization header
	req.Header.Set("Authorization", "invalid_token")
	rr = httptest.NewRecorder()
	authMiddleware.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusForbidden)
	}

	// Test with valid Authorization header
	req.Header.Set("Authorization", "Bearer valid_token")
	rr = httptest.NewRecorder()
	authMiddleware.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
