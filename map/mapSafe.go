package main

import "sync"

type MapRMutex struct {
	m map[string]interface{}
	sync.RWMutex
}

func NewMapRMutex() MapRMutex {
	return MapRMutex{
		m:       make(map[string]interface{}),
		RWMutex: sync.RWMutex{},
	}
}

func (m MapRMutex) Set(key string, value interface{}) {
	m.Lock()
	m.m[key] = value
	m.Unlock()
}

func (m MapRMutex) Get(key string) (interface{}, bool) {
	m.RLock()
	val, ok := m.m[key]
	m.RUnlock()
	return val, ok
}
