package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type DeviceStore interface {
	CreateDevice(device Device) error
	GetDevice(id string) (Device, error)
	DeleteDevice(id string) error
}

type DeviceStoreClient struct {
	baseUrl string
}

func NewDeviceStoreClient(baseUrl string) *DeviceStoreClient {
	return &DeviceStoreClient{baseUrl: baseUrl}
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func (c *DeviceStoreClient) CreateDevice(device Device) error {
	deviceJson, _ := json.Marshal(device)
	res, err := http.Post(c.baseUrl+"/devices", "application/json", bytes.NewBuffer(deviceJson))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return errors.New("unexpected status code")
	}
	return nil
}

func (c *DeviceStoreClient) GetDevice(id string) (Device, error) {
	res, err := http.Get(c.baseUrl + "/devices/" + id)
	if err != nil {
		return Device{}, err
	}
	if res.StatusCode != http.StatusOK {
		return Device{}, errors.New("unexpected status code")
	}
	var device Device
	err = json.NewDecoder(res.Body).Decode(&device)
	if err != nil {
		return Device{}, err
	}
	err = validate.Struct(device)
	if err != nil {
		return Device{}, err
	}
	return device, nil
}

func (c *DeviceStoreClient) DeleteDevice(id string) error {
	req, err := http.NewRequest("DELETE", c.baseUrl+"/devices/"+id, nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		return errors.New("unexpected status code")
	}
	return nil
}
