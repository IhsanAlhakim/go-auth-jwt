package handlers

import (
	"encoding/json"
	"net/http"
)

type Payload struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type P = Payload

func Response(w http.ResponseWriter, payload Payload, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(payload)
}
