package service

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	client "github.com/alexbavinton/home-automation/device-store/pkg/client"
	"github.com/gorilla/mux"
)

type mockGetDevicer struct {
	called     bool
	calledWith string
	returns    client.Device
	errorWith  error
}

func (m *mockGetDevicer) GetDevice(id string) (client.Device, error) {
	m.called = true
	m.calledWith = id
	return m.returns, m.errorWith
}

func TestGetDevice(t *testing.T) {
	t.Run("responds with 404 if device not found", func(t *testing.T) {
		mockGetDevice := mockGetDevicer{
			returns:   client.Device{},
			errorWith: errors.New("device not found"),
		}

		handler := getDeviceHandler(&mockGetDevice)

		req := httptest.NewRequest("GET", "/devices/1", nil)
		res := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": "1",
		})

		handler(res, req)

		if !mockGetDevice.called {
			t.Error("Expected GetDevice to be called")
		}

		if mockGetDevice.calledWith != "1" {
			t.Errorf("Expected GetDevice to be called with 1, got %s", mockGetDevice.calledWith)
		}

		if res.Code != http.StatusNotFound {
			t.Errorf("Expected status code %d, got %d", http.StatusNotFound, res.Code)
		}

	})

	t.Run("responds with 500 if another error occurs", func(t *testing.T) {
		mockGetDevice := mockGetDevicer{
			returns:   client.Device{},
			errorWith: errors.New("some error"),
		}

		handler := getDeviceHandler(&mockGetDevice)

		req := httptest.NewRequest("GET", "/devices/1", nil)
		res := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": "1",
		})

		handler(res, req)

		if !mockGetDevice.called {
			t.Error("Expected GetDevice to be called")
		}

		if mockGetDevice.calledWith != "1" {
			t.Errorf("Expected GetDevice to be called with 1, got %s", mockGetDevice.calledWith)
		}

		if res.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, res.Code)
		}
	})

	t.Run("responds with 200 and the device if found", func(t *testing.T) {
		expectedDevice := client.Device{
			ID:          "1",
			Name:        "bulb-1",
			Description: "a bulb",
			Type:        "bulb",
		}
		mockGetDevice := mockGetDevicer{
			returns:   expectedDevice,
			errorWith: nil,
		}

		handler := getDeviceHandler(&mockGetDevice)

		req := httptest.NewRequest("GET", "/devices/1", nil)
		res := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": "1",
		})

		handler(res, req)

		if !mockGetDevice.called {
			t.Error("Expected GetDevice to be called")
		}

		if mockGetDevice.calledWith != "1" {
			t.Errorf("Expected GetDevice to be called with 1, got %s", mockGetDevice.calledWith)
		}

		if res.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
		}

		// var device client.Device
		// json.NewDecoder(res.Body).Decode(&device)

		// if device != expectedDevice {
		// 	t.Errorf("Expected device to be %v, got %v", expectedDevice, device)
		// }

	})
}
