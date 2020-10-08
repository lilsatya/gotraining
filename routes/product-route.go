package routes

import (
	"gotraining/modules/product"
	"net/http"

	"github.com/gorilla/mux"
)

// ProductRoute route
func ProductRoute(r *mux.Router) {
	r.HandleFunc("/product", product.Create).Methods(http.MethodPost)
	r.HandleFunc("/product", product.List).Methods(http.MethodGet)
}
