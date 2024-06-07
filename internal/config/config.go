// Package config provides functionality to parse and manage application configuration.
package config

import (
	"flag"
	"fmt"

	"github.com/MrTomSawyer/go-kv-storage/internal/apperrors"
)

// Config represents the application configuration.
type Config struct {
	// Port is the server port.
	Port uint

	// TTL is the default storage TTL (Time-To-Live) for key-value pairs.
	TTL uint

	CleanFreq uint
}

// New creates a new Config instance.
func New() *Config {
	return &Config{}
}

// MustInit initializes the configuration from command-line flags. It panics if any required flag is not provided or has invalid value.
func (c *Config) MustInit() {
	flag.UintVar(&c.Port, "port", 8080, "server port")
	flag.UintVar(&c.TTL, "ttl", 5, "default storage ttl")
	flag.UintVar(&c.CleanFreq, "cf", 5, "storage clean frequency")
	flag.Parse()

	if c.Port < 1 || c.Port > 65535 {
		fmt.Println("Error: Port number is out of range (1-65535)")
		panic(apperrors.ErrWrongPort)
	}

	if c.TTL < 1 {
		fmt.Println("Error: TTL must not be less than 1")
		panic(apperrors.ErrWrongTTL)
	}

	if c.CleanFreq < 1 {
		fmt.Println("Error: Clean frequency must not be less than 1")
		panic(apperrors.ErrWrongCleanFreq)
	}
}
