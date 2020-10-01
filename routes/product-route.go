package routes

import (
	"gotraining/modules/product"
	"net/http"

	"github.com/gorilla/mux"
)

// ProductRoutes route
func ProductRoutes(r *mux.Router) {
	r.HandleFunc("/products", product.List).Methods(http.MethodGet)
}
