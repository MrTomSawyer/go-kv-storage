package service

import (
	"context"
	"github.com/MrTomSawyer/go-kv-storage/internal/models" // Importing the models package for the Entry type
	"time"                                                 // Importing the time package
)

// KVService provides methods to interact with the key-value storage.
type KVService struct {
	EntityProvider EntityProvider
	EntityModifier EntityModifier
}

// EntityProvider defines an interface for retrieving entries from the storage.
type EntityProvider interface {
	Get(ctx context.Context, key string) (models.Entry, bool)
}

// EntityModifier defines an interface for modifying entries in the storage.
type EntityModifier interface {
	Set(ctx context.Context, key string, value string, TTL ...time.Duration)
	Delete(ctx context.Context, key string)
}

// New creates a new instance of KVService.
func New(s EntityProvider, m EntityModifier) *KVService {
	return &KVService{
		EntityProvider: s,
		EntityModifier: m,
	}
}

// Get retrieves a value by key using the EntityProvider.
func (k *KVService) Get(ctx context.Context, key string) (models.Entry, bool) {
	return k.EntityProvider.Get(ctx, key)
}

// Set adds or updates a key-value pair with a specified TTL using the EntityModifier.
func (k *KVService) Set(ctx context.Context, key string, value string, TTL time.Duration) {
	k.EntityModifier.Set(ctx, key, value, TTL)
}

// Delete removes a value by key using the EntityModifier.
func (k *KVService) Delete(ctx context.Context, key string) {
	k.EntityModifier.Delete(ctx, key)
}
