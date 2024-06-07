package storage

import (
	"context"
	"github.com/MrTomSawyer/go-kv-storage/internal/daemon"
	"github.com/MrTomSawyer/go-kv-storage/internal/models"
	"time"
)

type UnsafeStorage struct {
	S map[string]models.Entry
}

func NewUnsafe(size uint) *UnsafeStorage {
	return &UnsafeStorage{
		S: make(map[string]models.Entry, size),
	}
}

func InitUnsafeStorage(ctx context.Context, size uint, cleanFreq time.Duration) *UnsafeStorage {
	s := NewUnsafe(size)
	daemon.CleanUnSafeStorage(ctx, s, cleanFreq*time.Second)

	return s
}

func (s *UnsafeStorage) Get(key string) (models.Entry, bool) {
	item, exists := s.S[key]
	if !exists || time.Now().After(item.ExpiresAt) {
		return models.Entry{}, false
	}

	return item, true
}

func (s *UnsafeStorage) Set(key string, value string, TTL time.Duration) {
	s.S[key] = models.Entry{
		Value:     value,
		ExpiresAt: time.Now().Add(TTL),
	}
}

func (s *UnsafeStorage) Delete(key string) {
	delete(s.S, key)
}
