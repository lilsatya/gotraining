package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Serve routes
func Serve() {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	StoreRoute(r)
	SupplierRoute(r)
	ProductRoute(r)

	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

	log.Println("ready to serve at localhost:8080")
}
