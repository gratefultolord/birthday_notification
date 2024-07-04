package tests

import (
	"birthday_notification/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSubscribe(t *testing.T) {
	req := httptest.NewRequest("POST", "http://example.com/subscribe", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.Subscribe)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Подписка успешна"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUnsubscribe(t *testing.T) {
	req := httptest.NewRequest("POST", "http://example.com/unsubscribe", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.Unsubscribe)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Отписка успешна"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
