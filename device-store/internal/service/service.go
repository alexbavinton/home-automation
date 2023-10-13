package service

import (
	"net/http"

	store "github.com/alexbavinton/home-automation/device-store/internal/store"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

// Run starts the device-store service
func Run(redisHost string, redisPort string) {
	conn, err := redis.Dial("tcp", redisHost+":"+redisPort)
	if err != nil {
		panic(err)
	}
	deviceStore := store.NewDeviceStore(conn)
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("Get")
	r.HandleFunc("/devices", createDeviceHandler(deviceStore)).Methods("POST")
	r.HandleFunc("/devices/{id}", getDeviceHandler(deviceStore)).Methods("GET")
	r.HandleFunc("/devices/{id}", deleteDeviceHandler(deviceStore)).Methods("DELETE")
	r.HandleFunc("/devices", getDevicesHandler(deviceStore)).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
