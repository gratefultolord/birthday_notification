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
	// Устанавливаем текущую дату для тестов
	today := time.Now().Format("2006-01-02")
	return []models.User{
		{ID: 1, Name: "Alice", Birthdate: today, Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Birthdate: "1988-07-05", Email: "bob@example.com"},
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

	expected := `["Alice"]`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
