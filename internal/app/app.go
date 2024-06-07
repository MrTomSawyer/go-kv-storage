package app

import (
	"context"
	grpcapp "github.com/MrTomSawyer/go-kv-storage/internal/app/grpc"
	"github.com/MrTomSawyer/go-kv-storage/internal/config"
	"github.com/MrTomSawyer/go-kv-storage/internal/service"
	"github.com/MrTomSawyer/go-kv-storage/internal/storage"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

// New creates a new gRPC app instance
func New(ctx context.Context, cfg *config.Config) *App {
	CleanFreq := time.Duration(cfg.CleanFreq)
	TTL := time.Duration(cfg.TTL)

	st, _ := storage.InitStorage(ctx, CleanFreq, TTL)
	kvService := service.New(st, st)

	grpcApp := grpcapp.New(cfg, kvService)

	return &App{
		GRPCServer: grpcApp,
	}
}
