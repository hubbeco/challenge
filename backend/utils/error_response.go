package utils

import (
	"encoding/json"
	"net/http"
)

// Compoem o corpo da resposta de erro
type ErrorResponse struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

func RespondWithError(w http.ResponseWriter, status int, title, detail, instance string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	errorResponse := ErrorResponse{
		Type:     "about:blank",
		Title:    title,
		Detail:   detail,
		Instance: instance,
	}

	json.NewEncoder(w).Encode(errorResponse)
}
