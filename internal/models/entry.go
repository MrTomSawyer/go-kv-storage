package models

import "time"

// Entry struct represents a map entry
type Entry struct {
	Value     string
	ExpiresAt time.Time
}
