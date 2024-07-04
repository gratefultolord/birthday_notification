package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	createTables()
	insertTestData()
}

func createTables() {
	createUsersTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        birthdate TEXT NOT NULL,
        email TEXT NOT NULL
    );`
	_, err := DB.Exec(createUsersTableSQL)
	if err != nil {
		log.Fatalf("Error creating users table: %q", err)
	}

	createSubscriptionsTableSQL := `
    CREATE TABLE IF NOT EXISTS subscriptions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        subscriber_email TEXT NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id)
    );`
	_, err = DB.Exec(createSubscriptionsTableSQL)
	if err != nil {
		log.Fatalf("Error creating subscriptions table: %q", err)
	}
}

func insertTestData() {
	testUsers := []struct {
		Name      string
		Birthdate string
		Email     string
	}{
		{"Райан", "1980-07-04", "gosling@example.com"},
		{"Дуэйн", "1972-07-04", "rock@example.com"},
		{"Квентин", "1963-07-04", "tarantino@example.com"},
	}

	for _, user := range testUsers {
		_, err := DB.Exec("INSERT INTO users (name, birthdate, email) VALUES (?, ?, ?)", user.Name, user.Birthdate, user.Email)
		if err != nil {
			log.Fatalf("Error inserting test data: %q", err)
		}
	}
}
