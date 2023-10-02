package store

import (
	types "github.com/alexbavinton/home-automation/device-store/pkg/types"
	"github.com/gomodule/redigo/redis"
)

type Store interface {
	CreateDevice(device types.Device) error
	GetDevice(id string) (types.Device, error)
	DeleteDevice(id string) error
}

type RedisStore struct {
	Handler handlerWrapper
}

func NewRedisStore(conn *redis.Conn) *RedisStore {
	handler := newWrappedHandler(conn)
	return &RedisStore{Handler: handler}
}
