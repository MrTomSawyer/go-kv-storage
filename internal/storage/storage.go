package storage

import (
	"context"
	"github.com/MrTomSawyer/go-kv-storage/internal/models"
	"log"
	"sync"
	"time"
)

// Storage represents a thread-safe key-value storage
type Storage struct {
	S          sync.Map
	DefaultTTL time.Duration
}

// New creates a new Storage instance
func New(TTL time.Duration) *Storage {
	return &Storage{DefaultTTL: TTL}
}

// InitStorage creates a new Storage instance and starts a cleanup process
func InitStorage(ctx context.Context, cleanFreq time.Duration, TTL time.Duration) (*Storage, error) {
	s := New(TTL)
	go CleanStorage(ctx, s, cleanFreq)

	return s, nil
}

// Get retrieves a value by key if it exists and has not expired
func (s *Storage) Get(_ context.Context, key string) (models.Entry, bool) {
	item, exists := s.S.Load(key)
	if !exists {
		return models.Entry{}, false
	}

	entry := item.(models.Entry)
	if time.Now().After(entry.ExpiresAt) {
		return models.Entry{}, false
	}

	log.Printf("value for key %s found \n", key)
	return entry, true
}

// Set adds or updates a key-value pair with a DefaultTTL (time-to-live)
func (s *Storage) Set(_ context.Context, key string, value string, TTL ...time.Duration) {
	finalTTL := s.DefaultTTL

	// set custom TTL if one is provided
	if len(TTL) > 0 {
		finalTTL = TTL[0]
	}

	entry := models.Entry{
		Value:     value,
		ExpiresAt: time.Now().Add(finalTTL),
	}

	log.Printf("key %s value %s are saved \n", key, value)
	s.S.Store(key, entry)
}

// Delete removes a value by the given key
func (s *Storage) Delete(_ context.Context, key string) {
	s.S.Delete(key)
	log.Printf("key %s deleted \n", key)
}
