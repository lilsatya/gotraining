package routes

import (
	"gotraining/modules/store"
	"net/http"

	"github.com/gorilla/mux"
)

// StoreRoute route
func StoreRoute(r *mux.Router) {
	r.HandleFunc("/store", store.Create).Methods(http.MethodPost)
	r.HandleFunc("/store", store.List).Methods(http.MethodGet)
	r.HandleFunc("/store/{id:[0-9]+}", store.Read).Methods(http.MethodGet)
	r.HandleFunc("/store/{id:[0-9]+}", store.Update).Methods(http.MethodPut)
	r.HandleFunc("/store/{id:[0-9]+}", store.Delete).Methods(http.MethodDelete)
}
