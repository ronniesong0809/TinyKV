package store

import (
	"errors"
	"sync"
	"time"
)

type KVStore struct {
	mu   sync.RWMutex
	data map[string]valueWithTTL
}

type valueWithTTL struct {
	Value      string
	Expiration time.Time
}

var store = KVStore{
	data: make(map[string]valueWithTTL),
}

func Get(key string) (string, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	value, exists := store.data[key]
	if !exists {
		return "", errors.New("key not found")
	}
	if value.Expiration.Before(time.Now()) {
		delete(store.data, key)
		return "", errors.New("key expired")
	}
	return value.Value, nil
}

func Set(key, value string, ttl time.Duration) {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.data[key] = valueWithTTL{
		Value:      value,
		Expiration: time.Now().Add(ttl),
	}
}

func Update(key, value string, ttl time.Duration) error {
	store.mu.Lock()
	defer store.mu.Unlock()
	if _, exists := store.data[key]; !exists {
		return errors.New("key not found")
	}
	store.data[key] = valueWithTTL{
		Value:      value,
		Expiration: time.Now().Add(ttl),
	}
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
