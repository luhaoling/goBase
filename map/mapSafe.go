package main

import "sync"

type MapRMutex struct {
	m map[string]interface{}
	sync.Mutex
}

func NewMapRMutex() MapRMutex {
	return MapRMutex{
		m: make(map[string]interface{}),
	}
}

func (m *MapRMutex) Set(key string, value interface{}) {
	m.Lock()
	m.m[key] = value
	m.Unlock()
}

func (m *MapRMutex) Get(key string) (interface{}, bool) {
	m.Lock()
	val, ok := m.m[key]
	m.Unlock()
	return val, ok
}
