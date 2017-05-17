package utils

import (
	"sync"

	"github.com/gorilla/websocket"
)

// WebsocketConnMap ...
type WebsocketConnMap struct {
	source map[string]*websocket.Conn
	lock   sync.RWMutex
}

// NewWebsocketConnMap ...
func NewWebsocketConnMap() *WebsocketConnMap {
	return &WebsocketConnMap{source: make(map[string]*websocket.Conn)}
}

// SetNX ...
func (m *WebsocketConnMap) SetNX(key string, value *websocket.Conn) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.source[key]; ok {
		return false
	}
	m.source[key] = value
	return true
}

// Del ...
func (m *WebsocketConnMap) Del(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.source, key)
}

// MGet ...
func (m *WebsocketConnMap) MGet(keys []string) []*websocket.Conn {
	m.lock.RLock()
	defer m.lock.RUnlock()
	result := []*websocket.Conn{}
	for _, key := range keys {
		value, ok := m.source[key]
		if ok {
			result = append(result, value)
		}
	}
	return result
}
