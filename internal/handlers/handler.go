package handlers

import (
	"encoding/json"
	"net/http"
	"notification_service/internal/services"
)

type Handler struct {
	srvc *services.Service
}

func NewHandler(srvc *services.Service) *Handler {
	return &Handler{srvc: srvc}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"message": "pong"}

		json.NewEncoder(w).Encode(response)
	})

	mux.HandleFunc("/event", h.HandleEvent)

	return mux
}
