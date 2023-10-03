package client

// Package client and an api client for the device-store service

// Device represents a registered IoT device
type Device struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Type        string `json:"type" validate:"required"`
}
