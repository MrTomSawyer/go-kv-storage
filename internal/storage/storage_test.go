package storage

import (
	"context"
	"testing"
	"time"
)

func TestUnsafeStorage(t *testing.T) {
	// Initialize UnsafeStorage with a clean frequency of 1 second for testing
	cleanFreq := 1 * time.Second
	TTL := 3 * time.Second
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize UnsafeStorage with cleaner
	s, err := InitStorage(ctx, cleanFreq, TTL)
	if err != nil {
		t.Fatalf("Error initializing storage: %v", err)
	}

	// Test Set and Get methods
	s.Set("key1", "value1", 2*time.Second)
	s.Set("key2", "value2", 5*time.Second)

	// Wait for some time to let cleaner do its job
	time.Sleep(3 * time.Second)

	// Check if expired entries are cleaned up
	if _, exists := s.Get("key1"); exists {
		t.Error("Expected key1 to be expired and cleaned up, but it still exists")
	}
	if _, exists := s.Get("key2"); !exists {
		t.Error("Expected key2 to exist, but it's cleaned up prematurely")
	}

	// Test Delete method
	s.Delete("key2")
	if _, exists := s.Get("key2"); exists {
		t.Error("Expected key2 to be deleted, but it still exists")
	}
}
