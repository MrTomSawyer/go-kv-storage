package storage

import (
	"context"
	"github.com/MrTomSawyer/go-kv-storage/internal/daemon"
	"github.com/MrTomSawyer/go-kv-storage/internal/models"
	"sync"
	"time"
)

// SafeStorage represents a thread-safe key-value storage
type SafeStorage struct {
	S sync.Map
}

// NewSafe New creates a new Storage instance
func NewSafe() *SafeStorage {
	return &SafeStorage{}
}

func InitSafeStorage(ctx context.Context, cleanFreq time.Duration) *SafeStorage {
	s := NewSafe()
	daemon.CleanSafeStorage(ctx, s, cleanFreq*time.Second)

	return s
}

// Get retrieves a value by key if it exists and has not expired
func (s *SafeStorage) Get(key string) (models.Entry, bool) {
	item, exists := s.S.Load(key)
	if !exists {
		return models.Entry{}, false
	}

	entry := item.(models.Entry)
	if time.Now().After(entry.ExpiresAt) {
		return models.Entry{}, false
	}

	return entry, true
}

// Set adds or updates a key-value pair with a TTL (time-to-live)
func (s *SafeStorage) Set(key string, value string, TTL time.Duration) {
	entry := models.Entry{
		Value:     value,
		ExpiresAt: time.Now().Add(TTL),
	}
	s.S.Store(key, entry)
}

// Delete removes a value by the given key
func (s *SafeStorage) Delete(key string) {
	s.S.Delete(key)
}
