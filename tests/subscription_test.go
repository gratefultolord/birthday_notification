package tests

import (
	"birthday_notification/internal/handlers"
	"birthday_notification/internal/repository"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	createTables(db)
	insertTestUsers(db)
	return db
}

func createTables(db *sql.DB) {
	createUsersTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        birthdate TEXT NOT NULL,
        email TEXT NOT NULL
    );`
	_, err := db.Exec(createUsersTableSQL)
	if err != nil {
		panic(err)
	}

	createSubscriptionsTableSQL := `
    CREATE TABLE IF NOT EXISTS subscriptions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        subscriber_email TEXT NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id)
    );`
	_, err = db.Exec(createSubscriptionsTableSQL)
	if err != nil {
		panic(err)
	}
}

func insertTestUsers(db *sql.DB) {
	testUsers := []struct {
		Name      string
		Birthdate string
		Email     string
	}{
		{"Alice", "1990-07-04", "alice@example.com"},
		{"Bob", "1988-07-05", "bob@example.com"},
		{"Charlie", "1992-07-04", "charlie@example.com"},
	}

	for _, user := range testUsers {
		_, err := db.Exec("INSERT INTO users (name, birthdate, email) VALUES (?, ?, ?)", user.Name, user.Birthdate, user.Email)
		if err != nil {
			panic(err)
		}
	}
}

func TestSubscribe(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewSubscriptionRepository(db)
	handler := handlers.NewSubscriptionHandler(repo)

	reqBody := `{"user_id": 1, "email": "subscriber@example.com"}`
	req := httptest.NewRequest("POST", "http://example.com/subscribe", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	httpHandler := http.HandlerFunc(handler.Subscribe)
	httpHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Подписка успешна"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUnsubscribe(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewSubscriptionRepository(db)
	handler := handlers.NewSubscriptionHandler(repo)

	_, err := db.Exec("INSERT INTO subscriptions (user_id, subscriber_email) VALUES (?, ?)", 1, "subscriber@example.com")
	if err != nil {
		t.Fatalf("Failed to insert subscription: %v", err)
	}

	reqBody := `{"user_id": 1, "email": "subscriber@example.com"}`
	req := httptest.NewRequest("POST", "http://example.com/unsubscribe", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	httpHandler := http.HandlerFunc(handler.Unsubscribe)
	httpHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Отписка успешна"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
