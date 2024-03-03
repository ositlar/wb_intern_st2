package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ositlar/go-http/internal/model"
)

func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "only Post method allowed",
		})
		return
	}
	updatedEvent, err := getEventForUpdateOrDelete(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "couldn't create event for update",
		})
		return
	}
	err = model.Cache.UpdateEvent(updatedEvent.ID, updatedEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "event was not found by ID",
		})
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]model.Event{
			"result": updatedEvent,
		})
	}
}
