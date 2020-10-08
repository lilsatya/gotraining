package product

import (
	"encoding/json"
	"gotraining/entities"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var p entities.Product
	handler := newHandler()
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		panic("Invalid request payload")
	}
	defer r.Body.Close()

	product, _ := handler.create(p)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// List will fetch all data
func List(w http.ResponseWriter, r *http.Request) {
	handler := newHandler()

	products, _ := handler.list()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
