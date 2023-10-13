package store

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	client "github.com/alexbavinton/home-automation/device-store/pkg/client"
	"github.com/rafaeljusto/redigomock"
)

func TestCreateDevice(t *testing.T) {
	t.Run("Creates device and adds device to sets", func(t *testing.T) {
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

}

func TestGetDevice(t *testing.T) {
	t.Run("Returns device if exists", func(t *testing.T) {
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

	t.Run("Returns error if device does not exist", func(t *testing.T) {
		conn := redigomock.NewConn()

		getDeviceCommand := conn.Command("JSON.GET", "device:1").Expect(nil)
		store := NewDeviceStore(conn)
		_, err := store.GetDevice("1")

		if err == nil {
			t.Errorf("Expected error %v, got %v", errors.New("device not found"), err)
		}

		if err.Error() != "device not found" {
			t.Errorf("Expected error %v, got %v", errors.New("device not found"), err)
		}

		if conn.Stats(getDeviceCommand) != 1 {
			t.Errorf("Expected 1 call to JSON.GET, got %d", conn.Stats(getDeviceCommand))
		}
	})
}

func TestDeleteDevice(t *testing.T) {
	t.Run("Deletes device and removes device from sets", func(t *testing.T) {
		conn := redigomock.NewConn()
		device := client.Device{
			ID:          "1",
			Name:        "bulb-1",
			Description: "a bulb",
			Type:        "bulb",
		}

		deviceJson, _ := json.Marshal(device)

		getDeviceCommand := conn.Command("JSON.GET", "device:1").Expect(string(deviceJson))
		deleteDeviceCommand := conn.Command("JSON.DEL", "device:1", ".").Expect("ok")
		removeFromTypeCommand := conn.Command("SREM", "device-type:bulb", "1").Expect("ok")
		removeFromDevicesCommand := conn.Command("SREM", "devices", "1").Expect("ok")
		store := NewDeviceStore(conn)
		err := store.DeleteDevice("1")

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if conn.Stats(getDeviceCommand) != 1 {
			t.Errorf("Expected 1 call to JSON.GET, got %d", conn.Stats(getDeviceCommand))
		}

		if conn.Stats(deleteDeviceCommand) != 1 {
			t.Errorf("Expected 1 call to DEL, got %d", conn.Stats(deleteDeviceCommand))
		}

		if conn.Stats(removeFromTypeCommand) != 1 {
			t.Errorf("Expected 1 call to SREM, got %d", conn.Stats(removeFromTypeCommand))
		}

		if conn.Stats(removeFromDevicesCommand) != 1 {
			t.Errorf("Expected 1 call to SREM, got %d", conn.Stats(removeFromDevicesCommand))
		}
	})
	t.Run("Returns error if device does not exist", func(t *testing.T) {
		conn := redigomock.NewConn()

		getDeviceCommand := conn.Command("JSON.GET", "device:1").Expect(nil)
		store := NewDeviceStore(conn)
		err := store.DeleteDevice("1")

		if err == nil {
			t.Errorf("Expected error %v, got %v", errors.New("device not found"), err)
		}

		if err.Error() != "device not found" {
			t.Errorf("Expected error %v, got %v", errors.New("device not found"), err)
		}

		if conn.Stats(getDeviceCommand) != 1 {
			t.Errorf("Expected 1 call to JSON.GET, got %d", conn.Stats(getDeviceCommand))
		}
	})
}

func TestGetDevices(t *testing.T) {
	t.Run("Returns all devices", func(t *testing.T) {
		conn := redigomock.NewConn()
		device1 := client.Device{
			ID:          "1",
			Name:        "bulb-1",
			Description: "a bulb",
			Type:        "bulb",
		}
		device2 := client.Device{
			ID:          "2",
			Name:        "bulb-2",
			Description: "a bulb",
			Type:        "bulb",
		}
		want := []client.Device{device1, device2}

		device1Json, _ := json.Marshal(device1)
		device2Json, _ := json.Marshal(device2)

		store := NewDeviceStore(conn)

		getDevicesCommand := conn.Command("SMEMBERS", "devices").Expect([]interface{}{"1", "2"})
		getDevice1Command := conn.Command("JSON.GET", "device:1").Expect(device1Json)
		getDevice2Command := conn.Command("JSON.GET", "device:2").Expect(device2Json)

		got, err := store.GetDevices()

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if conn.Stats(getDevicesCommand) != 1 {
			t.Errorf("Expected 1 call to JSON.GET, got %d", conn.Stats(getDevicesCommand))
		}

		if conn.Stats(getDevice1Command) != 1 {
			t.Errorf("Expected 1 call to JSON.GET, got %d", conn.Stats(getDevice1Command))
		}

		if conn.Stats(getDevice2Command) != 1 {
			t.Errorf("Expected 1 call to JSON.GET, got %d", conn.Stats(getDevice2Command))
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})
}
