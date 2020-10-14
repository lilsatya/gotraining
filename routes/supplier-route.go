package routes

import (
	"gotraining/modules/supplier"
	"net/http"

	"github.com/gorilla/mux"
)

// SupplierRoute route
func SupplierRoute(r *mux.Router) {
	r.HandleFunc("/supplier", supplier.Create).Methods(http.MethodPost)
	r.HandleFunc("/supplier", supplier.List).Methods(http.MethodGet)
	r.HandleFunc("/supplier/{id:[0-9]+}", supplier.Read).Methods(http.MethodGet)
	r.HandleFunc("/supplier/{id:[0-9]+}", supplier.Update).Methods(http.MethodPut)
	r.HandleFunc("/supplier/{id:[0-9]+}", supplier.Delete).Methods(http.MethodDelete)
}
