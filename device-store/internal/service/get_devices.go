package service

import (
	"encoding/json"
	"net/http"

	client "github.com/alexbavinton/home-automation/device-store/pkg/client"
)

type getDeviceser interface {
	GetDevices() ([]client.Device, error)
}

// GetDevicesHandler returns a handler for getting all devices
func getDevicesHandler(getDevices getDeviceser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		devices, err := getDevices.GetDevices()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(devices)
	}

}
