// Package server provides functionality to handle multiple client connections.
package server

import (
	"context"
	"fmt"
	"github.com/MrTomSawyer/go-kv-storage/internal/config"
	"github.com/MrTomSawyer/go-kv-storage/internal/storage"
	"net"
)

// Server represents a TCP server.
type Server struct {
	Cfg      *config.Config
	Storage  *storage.Storage
	listener net.Listener
}

func New(cfg *config.Config, storage *storage.Storage) *Server {
	return &Server{Cfg: cfg, Storage: storage}
}

func (s *Server) Start(ctx context.Context) error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Cfg.Port))
	if err != nil {
		panic(err)
	}

	s.listener = ln

	fmt.Println("Server started at port:", s.Cfg.Port)

	return nil
}
