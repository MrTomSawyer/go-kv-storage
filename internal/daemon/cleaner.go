package daemon

import (
	"context"
	"github.com/MrTomSawyer/go-kv-storage/internal/models"
	"github.com/MrTomSawyer/go-kv-storage/internal/storage"
	"time"
)

func CleanSafeStorage(ctx context.Context, m *storage.SafeStorage, cleanFreq time.Duration) {
	t := time.NewTicker(cleanFreq)
	defer t.Stop()

	IsCleanUpInProgress := false

	for {
		select {
		case <-ctx.Done():
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

func CleanUnSafeStorage(ctx context.Context, m *storage.UnsafeStorage, cleanFreq time.Duration) {
	t := time.NewTicker(cleanFreq)
	defer t.Stop()

	IsCleanUpInProgress := false

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			// to prevent possible overlaps, skip cleanup is one is already in progress
			if IsCleanUpInProgress {
				continue
			}

			IsCleanUpInProgress = true
			now := time.Now()

			for key, entry := range m.S {
				if now.After(entry.ExpiresAt) {
					delete(m.S, key)
				}
			}

			IsCleanUpInProgress = false
		}
	}
}
