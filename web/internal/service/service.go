package service

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("Get")
	r.HandleFunc("/devices", getDevicesPage()).Methods("GET")
	http.Handle("/", r)
	log.Print("Starting server on port 8080")
	log.Fatal((http.ListenAndServe(":8080", nil)))
}
