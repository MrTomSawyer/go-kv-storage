package models

import "time"

type Entry struct {
	Value     string
	ExpiresAt time.Time
}
