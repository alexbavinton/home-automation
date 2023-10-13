package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type mockDeleteDevicer struct {
	called     bool
	calledWith string
	errorsWith error
}

func (m *mockDeleteDevicer) DeleteDevice(id string) error {
	m.called = true
	m.calledWith = id
	return m.errorsWith
}

func TestDeleteDevice(t *testing.T) {
	t.Run("Responds with 204 if device exists", func(t *testing.T) {
		mockDeleteDevice := mockDeleteDevicer{}

		handler := deleteDeviceHandler(&mockDeleteDevice)

		req := httptest.NewRequest("DELETE", "/devices/1", nil)
		res := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": "1",
		})

		handler(res, req)

		if res.Code != http.StatusNoContent {
			t.Errorf("Expected status code %d, got %d", http.StatusNoContent, res.Code)
		}

		if !mockDeleteDevice.called {
			t.Errorf("Expected DeleteDevice to be called")
		}

		if mockDeleteDevice.calledWith != "1" {
			t.Errorf("Expected DeleteDevice to be called with %s, got %s", "1", mockDeleteDevice.calledWith)
		}
	})
}
