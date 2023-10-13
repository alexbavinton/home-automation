package service

import (
	"encoding/json"
	"net/http"

	client "github.com/alexbavinton/home-automation/device-store/pkg/client"
	"github.com/go-playground/validator/v10"
)

// Package handlers provides http handlers for the device-store service

// createDevicer is an interface for creating devices
type createDevicer interface {
	CreateDevice(device client.Device) error
}

// createDeviceHandler returns an http handler for creating devices
func createDeviceHandler(store createDevicer) http.HandlerFunc {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return func(w http.ResponseWriter, r *http.Request) {
		var device client.Device
		json.NewDecoder(r.Body).Decode(&device)
		err := validate.Struct(device)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = store.CreateDevice(device)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
