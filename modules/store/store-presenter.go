package store

import (
	"encoding/json"
	"net/http"
)

// List will fetch all data
func List(w http.ResponseWriter, r *http.Request) {
	handler := newHandler()

	stores, _ := handler.list()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stores)
}
