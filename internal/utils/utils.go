package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func ParseDate(dateStr string) (time.Time, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	log.Printf("Parsed date: %s -> %s\n", dateStr, t)
	return t, nil
}

func IsToday(date time.Time) bool {
	now := time.Now()
	location := now.Location()
	dateInLocation := date.In(location)

	log.Printf("Comparing dates: %s (given) vs %s (now)\n", dateInLocation.Format("2006-01-02"), now.Format("2006-01-02"))
	return now.Month() == dateInLocation.Month() && now.Day() == dateInLocation.Day()
}

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

func RespondWithError(w http.ResponseWriter, status int, message string) {
	RespondWithJSON(w, status, map[string]string{"error": message})
}
