package tests

import (
	"birthday_notification/internal/handlers"
	"birthday_notification/internal/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) GetAllUsers() ([]models.User, error) {
	today := time.Now().Format("2006-01-02")
	return []models.User{
		{ID: 1, Name: "Райан", Birthdate: today, Email: "gosling@example.com"},
		{ID: 2, Name: "Дуэйн", Birthdate: "1972-07-05", Email: "rock@example.com"},
	}, nil
}

func TestGetBirthdays(t *testing.T) {
	mockRepo := &MockUserRepository{}
	handler := handlers.NewBirthdayHandler(mockRepo)

	req := httptest.NewRequest("GET", "http://example.com/birthdays", nil)
	rr := httptest.NewRecorder()

	httpHandler := http.HandlerFunc(handler.GetBirthdays)
	httpHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `["Райан"]`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
