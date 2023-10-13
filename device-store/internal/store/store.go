package store

// Package store provides logic for accessing and modifying device data in redis

import (
	"encoding/json"
	"errors"

	client "github.com/alexbavinton/home-automation/device-store/pkg/client"
	"github.com/gomodule/redigo/redis"
)

// Store is an interface for accessing device data that is stored in redis
type Store interface {
	CreateDevice(device client.Device) error
	GetDevice(id string) (client.Device, error)
	DeleteDevice(id string) error
	GetDevicesByType(deviceType string) ([]client.Device, error)
	GetDevices() ([]client.Device, error)
	GetDeviceTypes() ([]string, error)
	Close() error
}

// DeviceStore is a redis implementation of the Store interface
type DeviceStore struct {
	conn redis.Conn
}

// NewDeviceStore returns a new DeviceStore
func NewDeviceStore(conn redis.Conn) *DeviceStore {
	return &DeviceStore{conn: conn}
}

// CreateDevice creates a new device in redis
func (s *DeviceStore) CreateDevice(device client.Device) error {
	deviceJson, err := json.Marshal(device)
	if err != nil {
		return err
	}
	s.conn.Send("JSON.SET", "device:"+device.ID, ".", string(deviceJson))
	s.conn.Send("SADD", "devices", device.ID)
	s.conn.Send("SADD", "device-types", device.Type)
	s.conn.Send("SADD", "device-type:"+device.Type, device.ID)
	return s.conn.Flush()
}

// GetDevice retrieves a device from redis
func (s *DeviceStore) GetDevice(id string) (client.Device, error) {
	deviceJson, err := redis.Bytes(s.conn.Do("JSON.GET", "device:"+id))
	if err != nil {
		return client.Device{}, errors.New("device not found")
	}

	var device client.Device
	err = json.Unmarshal(deviceJson, &device)
	if err != nil {
		return client.Device{}, err
	}
	return device, nil
}

// DeleteDevice deletes a device from redis
func (s *DeviceStore) DeleteDevice(id string) error {
	device, err := s.GetDevice(id)
	if err != nil {
		return err
	}
	s.conn.Send("JSON.DEL", "device:"+id, ".")
	s.conn.Send("SREM", "devices", id)
	s.conn.Send("SREM", "device-type:"+device.Type, id)
	return s.conn.Flush()
}

// GetDevices retrieves all devices from redis
func (s *DeviceStore) GetDevices() ([]client.Device, error) {
	deviceIds, err := redis.Strings(s.conn.Do("SMEMBERS", "devices"))
	if err != nil {
		return []client.Device{}, err
	}
	for _, id := range deviceIds {
		s.conn.Send("JSON.GET", "device:"+id)
	}
	deviceJsons, err := redis.ByteSlices(s.conn.Do(""))
	if err != nil {
		if err == redis.ErrNil {
			return []client.Device{}, nil
		}
		return []client.Device{}, err
	}
	var devices []client.Device
	for _, deviceJson := range deviceJsons {
		var device client.Device
		err = json.Unmarshal(deviceJson, &device)
		if err != nil {
			return []client.Device{}, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}

// GetDevicesByType retrieves all devices of a given type from redis
func (s *DeviceStore) GetDevicesByType(deviceType string) ([]client.Device, error) {
	panic("not implemented")
}

// GetDeviceTypes retrieves all device types from redis
func (s *DeviceStore) GetDeviceTypes() ([]string, error) {
	panic("not implemented")
}

// Close closes the redis connection
func (s *DeviceStore) Close() error {
	return s.conn.Close()
}
