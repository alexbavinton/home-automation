package store

import (
	"encoding/json"
	"reflect"
	"testing"

	client "github.com/alexbavinton/home-automation/device-store/pkg/client"
	"github.com/rafaeljusto/redigomock"
)

func TestDeviceStore(t *testing.T) {
	t.Run("CreateDevice", func(t *testing.T) {
		conn := redigomock.NewConn()
		device := client.Device{
			ID:          "1",
			Name:        "bulb-1",
			Description: "a bulb",
			Type:        "bulb",
		}

		deviceJson, _ := json.Marshal(device)

		createDeviceCommand := conn.Command("JSON.SET", "device:1", ".", string(deviceJson)).Expect("ok")
		addTypeCommand := conn.Command("SADD", "device-types", "bulb").Expect("ok")
		addToDeviceSetCommand := conn.Command("SADD", "devices", "1").Expect("ok")
		addDeviceToTypeCommand := conn.Command("SADD", "device-type:bulb", "1").Expect("ok")
		store := NewDeviceStore(conn)
		err := store.CreateDevice(device)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if conn.Stats(createDeviceCommand) != 1 {
			t.Errorf("Expected 1 call to JSON.SET, got %d", conn.Stats(createDeviceCommand))
		}

		if conn.Stats(addTypeCommand) != 1 {
			t.Errorf("Expected 1 type to be added to set, got %d", conn.Stats(addTypeCommand))
		}

		if conn.Stats(addToDeviceSetCommand) != 1 {
			t.Errorf("Expected 1 device to be added to set, got %d", conn.Stats(addToDeviceSetCommand))
		}

		if conn.Stats(addDeviceToTypeCommand) != 1 {
			t.Errorf("Expected 1 device to be added to type, got %d", conn.Stats(addDeviceToTypeCommand))
		}
	})

	t.Run("GetDevice", func(t *testing.T) {
		conn := redigomock.NewConn()
		want := client.Device{
			ID:          "1",
			Name:        "bulb-1",
			Description: "a bulb",
			Type:        "bulb",
		}

		deviceJson, _ := json.Marshal(want)

		getDeviceCommand := conn.Command("JSON.GET", "device:1").Expect(string(deviceJson))
		store := NewDeviceStore(conn)
		got, err := store.GetDevice("1")

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if conn.Stats(getDeviceCommand) != 1 {
			t.Errorf("Expected 1 call to JSON.GET, got %d", conn.Stats(getDeviceCommand))
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})
}
