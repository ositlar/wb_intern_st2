package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ositlar/go-http/internal/model"
)

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Wrong method : only Post is allowed",
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
	err = model.Cache.DeleteEvent(updatedEvent.ID, updatedEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "event was not found by ID",
		})
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"result": "deleted",
		})
	}
}
