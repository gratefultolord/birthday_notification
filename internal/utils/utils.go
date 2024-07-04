package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

// ParseDate parses a string date in the format "YYYY-MM-DD" and returns a time.Time object.
func ParseDate(dateStr string) (time.Time, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// IsToday checks if the given date is today.
func IsToday(date time.Time) bool {
	now := time.Now()
	return now.Year() == date.Year() && now.YearDay() == date.YearDay()
}

// RespondWithJSON sends a JSON response with the given status code and payload.
func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// RespondWithError sends a JSON response with the given status code and error message.
func RespondWithError(w http.ResponseWriter, status int, message string) {
	RespondWithJSON(w, status, map[string]string{"error": message})
}
