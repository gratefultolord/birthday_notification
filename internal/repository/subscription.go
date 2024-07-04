package repository

import (
	"database/sql"
)

type SubscriptionRepository struct {
	DB *sql.DB
}

func NewSubscriptionRepository(db *sql.DB) *SubscriptionRepository {
	return &SubscriptionRepository{DB: db}
}

func (repo *SubscriptionRepository) Subscribe(userID int, subscriberEmail string) error {
	_, err := repo.DB.Exec("INSERT INTO subscriptions (user_id, subscriber_email) VALUES (?, ?)", userID, subscriberEmail)
	return err
}

func (repo *SubscriptionRepository) Unsubscribe(userID int, subscriberEmail string) error {
	_, err := repo.DB.Exec("DELETE FROM subscriptions WHERE user_id = ? AND subscriber_email = ?", userID, subscriberEmail)
	return err
}

func (repo *SubscriptionRepository) GetSubscriptions(subscriberEmail string) ([]int, error) {
	rows, err := repo.DB.Query("SELECT user_id FROM subscriptions WHERE subscriber_email = ?", subscriberEmail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userIDs []int
	for rows.Next() {
		var userID int
		if err := rows.Scan(&userID); err != nil {
			return nil, err
		}
		userIDs = append(userIDs, userID)
	}
	return userIDs, nil
}
