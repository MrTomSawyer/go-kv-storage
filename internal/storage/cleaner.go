package storage

import (
	"context"
	"github.com/MrTomSawyer/go-kv-storage/internal/models"
	"log"
	"time"
)

// CleanStorage periodically cleans up expired entries in the storage
func CleanStorage(ctx context.Context, m *Storage, cleanFreq time.Duration) {
	t := time.NewTicker(cleanFreq)
	defer t.Stop()

	IsCleanUpInProgress := false

	log.Printf("starting storage creanup")
	for {
		select {
		case <-ctx.Done():
			log.Printf("storage creanup stopped")
			return
		case <-t.C:
			// to prevent possible overlaps, skip cleanup is one is already in progress
			if IsCleanUpInProgress {
				continue
			}

			IsCleanUpInProgress = true
			now := time.Now()

			m.S.Range(func(key, value interface{}) bool {
				entry, ok := value.(models.Entry)
				if !ok {
					return true
				}
				if now.After(entry.ExpiresAt) {
					m.S.Delete(key)
				}
				return true
			})

			IsCleanUpInProgress = false
		}
	}
}
