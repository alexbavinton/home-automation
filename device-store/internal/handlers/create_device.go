package handlers

import (
	"encoding/json"
	"net/http"

	client "github.com/alexbavinton/home-automation/device-store/pkg/client"
	"github.com/go-playground/validator/v10"
)

// Package handlers provides http handlers for the device-store service

// CreateDevicer is an interface for creating devices
type CreateDevicer interface {
	CreateDevice(device client.Device) error
}

// CreateDeviceHandler returns an http handler for creating devices
func CreateDeviceHandler(store CreateDevicer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var device client.Device
		err := json.NewDecoder(r.Body).Decode(&device)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		validate := validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(device)
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
