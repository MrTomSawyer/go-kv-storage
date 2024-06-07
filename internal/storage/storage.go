package storage

import (
	"context"
	"github.com/MrTomSawyer/go-kv-storage/internal/models"
	"sync"
	"time"
)

// Storage represents a thread-safe key-value storage
type Storage struct {
	S          sync.Map
	defaultTTL time.Duration
}

// New creates a new Storage instance
func New(TTL time.Duration) *Storage {
	return &Storage{defaultTTL: TTL}
}

// InitStorage creates a new Storage instance and starts a cleanup process
func InitStorage(ctx context.Context, cleanFreq time.Duration, TTL time.Duration) (*Storage, error) {
	s := New(TTL)
	go CleanStorage(ctx, s, cleanFreq)

	return s, nil
}

// Get retrieves a value by key if it exists and has not expired
func (s *Storage) Get(key string) (models.Entry, bool) {
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

// Set adds or updates a key-value pair with a defaultTTL (time-to-live)
func (s *Storage) Set(key string, value string, TTL ...time.Duration) {
	finalTTL := s.defaultTTL

	// set custom TTL if one is provided
	if len(TTL) > 0 {
		finalTTL = TTL[0]
	}

	entry := models.Entry{
		Value:     value,
		ExpiresAt: time.Now().Add(finalTTL),
	}
	s.S.Store(key, entry)
}

// Delete removes a value by the given key
func (s *Storage) Delete(key string) {
	s.S.Delete(key)
}
