package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/ositlar/go-http/internal/model"
)

type inputData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Start       string `json:"start"`
	End         string `json:"end"`
}

func getEventForUpdateOrDelete(r *http.Request) (model.Event, error) {
	var updatedEvent model.Event
	err := json.NewDecoder(r.Body).Decode(&updatedEvent)
	if err != nil {
		return model.Event{}, err
	}

	return updatedEvent, nil
}

func GetDateQuery(r *http.Request) (time.Time, error) {
	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		return time.Time{}, errors.New("query parameter missing")
	}
	return time.Parse("2006-01-02", dateStr)
}
