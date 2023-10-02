package store

import (
	types "github.com/alexbavinton/home-automation/device-store/pkg/types"
)

type Store interface {
	CreateDevice(device types.Device) error
	GetDevice(id string) (types.Device, error)
	DeleteDevice(id string) error
}
