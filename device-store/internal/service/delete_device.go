package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

type deleteDevicer interface {
	DeleteDevice(id string) error
}

// deleteDeviceHandler returns an http handler for deleting devices
func deleteDeviceHandler(store deleteDevicer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		deviceID := vars["id"]

		err := store.DeleteDevice(deviceID)
		if err != nil {
			if err.Error() == "device not found" {
				http.NotFound(w, r)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
