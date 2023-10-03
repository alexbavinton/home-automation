package types

// Package types and an api client for the device-store service

// Device represents a registered IoT device
type Device struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}
