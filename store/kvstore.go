package store

import (
	"errors"
	"sync"
)

type KVStore struct {
	mu   sync.RWMutex
	data map[string]string
}

var store = KVStore{
	data: make(map[string]string),
}

func Get(key string) (string, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	value, exists := store.data[key]
	if !exists {
		return "", errors.New("key not found")
	}
	return value, nil
}

func Set(key, value string) {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.data[key] = value
}

func Update(key, value string) error {
	store.mu.Lock()
	defer store.mu.Unlock()
	if _, exists := store.data[key]; !exists {
		return errors.New("key not found")
	}
	store.data[key] = value
	return nil
}

func Delete(key string) error {
	store.mu.Lock()
	defer store.mu.Unlock()
	if _, exists := store.data[key]; !exists {
		return errors.New("key not found")
	}
	delete(store.data, key)
	return nil
}
