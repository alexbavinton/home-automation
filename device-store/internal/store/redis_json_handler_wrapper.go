package store

import (
	"github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson/v4"
)

type handlerWrapper interface {
	jsonSet(key, path string, obj interface{}) (res interface{}, err error)
	jsonGet(key, path string) (res interface{}, err error)
	jsonDel(key, path string) (res interface{}, err error)
}

type wrappedHandler struct {
	Handler *rejson.Handler
}

func (h *wrappedHandler) jsonSet(key, path string, obj interface{}) (res interface{}, err error) {
	return h.Handler.JSONSet(key, path, obj)
}

func (h *wrappedHandler) jsonGet(key, path string) (res interface{}, err error) {
	return h.Handler.JSONGet(key, path)
}

func (h *wrappedHandler) jsonDel(key, path string) (res interface{}, err error) {
	return h.Handler.JSONDel(key, path)
}

func newWrappedHandler(conn *redis.Conn) *wrappedHandler {
	handler := rejson.NewReJSONHandler()
	handler.SetRedigoClient(*conn)
	wrappedHandler := wrappedHandler{Handler: handler}
	return &wrappedHandler
}
