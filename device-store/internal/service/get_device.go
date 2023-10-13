package service

import (
	"encoding/json"
	"net/http"

	"github.com/alexbavinton/home-automation/device-store/pkg/client"
	"github.com/gorilla/mux"
)

type getDevicer interface {
	GetDevice(id string) (client.Device, error)
}

// getDeviceHandler returns an http handler for getting devices
func getDeviceHandler(store getDevicer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		deviceID := vars["id"]

		device, err := store.GetDevice(deviceID)
		if err != nil {
			if err.Error() == "device not found" {
				http.NotFound(w, r)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(device)

	}
}
