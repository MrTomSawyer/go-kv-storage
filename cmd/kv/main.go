package main

import (
	"context"
	"github.com/MrTomSawyer/go-kv-storage/internal/config"
	"github.com/MrTomSawyer/go-kv-storage/internal/server"
	"github.com/MrTomSawyer/go-kv-storage/internal/storage"
	"time"
)

func main() {
	cfg := config.New()
	cfg.MustInit()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cleanFreq := time.Duration(cfg.CleanFreq) * time.Second
	TTL := time.Duration(cfg.TTL) * time.Second

	st, err := storage.InitStorage(ctx, cleanFreq, TTL)
	if err != nil {
		panic(err)
	}

	s := server.New(cfg, st)

	err = s.Start(ctx)
	if err != nil {
		panic(err)
	}
}
