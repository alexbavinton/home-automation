package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/alexbavinton/home-automation/device-store/pkg/client"
)

type MockStore struct {
	called     bool
	calledWith client.Device
}

func (m *MockStore) CreateDevice(device client.Device) error {
	m.called = true
	m.calledWith = device
	return nil
}

func TestCreateDeviceHandler(t *testing.T) {
	t.Run("Responds with 400 if device is invalid shape", func(t *testing.T) {
		mockCreateDevice := MockStore{}

		handler := CreateDeviceHandler(&mockCreateDevice)

		req := httptest.NewRequest("PUT", "/device", bytes.NewBuffer([]byte(`{"name": "bar"}`)))
		res := httptest.NewRecorder()

		handler(res, req)

		if res.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, res.Code)
		}

	})

	t.Run("Calls CreateDevice on store with device", func(t *testing.T) {
		mockCreateDevice := MockStore{}

		handler := CreateDeviceHandler(&mockCreateDevice)

		device := client.Device{
			ID:          "1",
			Name:        "bulb-1",
			Description: "a bulb",
			Type:        "bulb",
		}

		deviceJson, _ := json.Marshal(device)

		req := httptest.NewRequest("PUT", "/devices", bytes.NewBuffer(deviceJson))
		res := httptest.NewRecorder()

		handler(res, req)

		if !mockCreateDevice.called {
			t.Errorf("Expected CreateDevice to be called")
		}

		if !reflect.DeepEqual(mockCreateDevice.calledWith, device) {
			t.Errorf("Expected CreateDevice to be called with %v, got %v", device, mockCreateDevice.calledWith)
		}

		if res.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, res.Code)
		}

	})
}
