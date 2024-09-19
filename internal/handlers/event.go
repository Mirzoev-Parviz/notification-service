package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"notification_service/internal/models"
)

func (h *Handler) HandleEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	var event models.Event

	if err := json.Unmarshal(body, &event); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing JSON: %v", err), http.StatusBadRequest)
		return
	}

	if err := h.srvc.SaveEvent(event); err != nil {
		http.Error(w, fmt.Sprintf("Error occured while wanted to save event: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
