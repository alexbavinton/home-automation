package store

import (
	"testing"

	client "github.com/alexbavinton/home-automation/device-store/pkg/client"
	"github.com/rafaeljusto/redigomock"
)

func TestDeviceStore(t *testing.T) {
	t.Run("CreateDevice", func(t *testing.T) {
		conn := redigomock.NewConn()
		createDeviceCommand := conn.Command("JSON.SET", "device:1", ".", `{"id":"1","name":"bulb-1","description":"a bulb","type":"bulb"}`).Expect("ok")
		addGroupCommand := conn.Command("SAAD", "device-types", "bulb").Expect("ok")
		addToDeviceSetCommand := conn.Command("SADD", "devices", "1").Expect("ok")
		addDeviceToGroupCommand := conn.Command("SADD", "device-type:bulb", "1").Expect("ok")
		store := NewDeviceStore(conn)
		err := store.CreateDevice(client.Device{
			ID:          "1",
			Name:        "bulb-1",
			Description: "a bulb",
			Type:        "bulb",
		})

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if conn.Stats(createDeviceCommand) != 1 {
			t.Errorf("Expected 1 call to JSON.SET, got %d", conn.Stats(createDeviceCommand))
		}

		if conn.Stats(addGroupCommand) != 1 {
			t.Errorf("Expected 1 group to be added to set, got %d", conn.Stats(addGroupCommand))
		}

		if conn.Stats(addToDeviceSetCommand) != 1 {
			t.Errorf("Expected 1 device to be added to set, got %d", conn.Stats(addToDeviceSetCommand))
		}

		if conn.Stats(addDeviceToGroupCommand) != 1 {
			t.Errorf("Expected 1 device to be added to group, got %d", conn.Stats(addDeviceToGroupCommand))
		}

	})
}
