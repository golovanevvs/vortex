package main

import "sync"

type SafeMap[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{m: make(map[K]V)}
}

func (sm *SafeMap[K, V]) Load(key K) (value V, ok bool) {
	sm.mu.RLock()
	value, ok = sm.m[key]
	sm.mu.RUnlock()
	return
}

func (sm *SafeMap[K, V]) Store(key K, value V) {
	sm.mu.Lock()
	sm.m[key] = value
	sm.mu.Unlock()
}

func (sm *SafeMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	sm.mu.Lock()
	if v, ok := sm.m[key]; ok {
		sm.mu.Unlock()
		return v, true
	}
	sm.m[key] = value
	sm.mu.Unlock()
	return value, false
}

func (sm *SafeMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	sm.mu.Lock()
	value, loaded = sm.m[key]
	if loaded {
		delete(sm.m, key)
	}
	sm.mu.Unlock()
	return
}

func (sm *SafeMap[K, V]) Delete(key K) {
	sm.mu.Lock()
	delete(sm.m, key)
	sm.mu.Unlock()
}

func (sm *SafeMap[K, V]) Range(f func(K, V) bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	for k, v := range sm.m {
		if !f(k, v) {
			break
		}
	}
}