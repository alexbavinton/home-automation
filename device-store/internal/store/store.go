package store

// Package store provides logic for accessing and modifying device data in redis

import (
	client "github.com/alexbavinton/home-automation/device-store/pkg/client"
	"github.com/gomodule/redigo/redis"
)

// Store is an interface for accessing device data that is stored in redis
type Store interface {
	CreateDevice(device client.Device) error
	GetDevice(id string) (client.Device, error)
	DeleteDevice(id string) error
}

// DeviceStore is a redis implementation of the Store interface
type DeviceStore struct {
	conn *redis.Conn
}

// NewDeviceStore returns a new DeviceStore
func NewDeviceStore(conn *redis.Conn) *DeviceStore {
	return &DeviceStore{conn: conn}
}

// CreateDevice creates a new device in redis
func (s *DeviceStore) CreateDevice(device client.Device) error {
	panic("not implemented")
}

// GetDevice retrieves a device from redis
func (s *DeviceStore) GetDevice(id string) (client.Device, error) {
	panic("not implemented")
}

// DeleteDevice deletes a device from redis
func (s *DeviceStore) DeleteDevice(id string) error {
	panic("not implemented")
}
