package service

import (
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/alexbavinton/home-automation/device-store/pkg/client"
)

type mockGetDevices struct {
	called  bool
	returns []client.Device
}

func (m *mockGetDevices) GetDevices() ([]client.Device, error) {
	m.called = true
	return m.returns, nil
}

func TestGetDevices(t *testing.T) {
	t.Run("Returns all devices", func(t *testing.T) {
		want := []client.Device{
			{
				ID:          "1",
				Name:        "bulb-1",
				Description: "a bulb",
				Type:        "bulb",
			},
			{
				ID:          "2",
				Name:        "bulb-2",
				Description: "a bulb",
				Type:        "bulb",
			},
		}
		mockStore := mockGetDevices{
			returns: want,
		}

		handler := getDevicesHandler(&mockStore)
		req := httptest.NewRequest("GET", "/devices", nil)
		res := httptest.NewRecorder()

		handler(res, req)

		if !mockStore.called {
			t.Error("Expected GetDevices to be called")
		}

		got := []client.Device{}
		json.NewDecoder(res.Body).Decode(&got)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}

	})
}
