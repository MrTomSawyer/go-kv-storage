package grpc

import (
	"fmt"
	"github.com/MrTomSawyer/go-kv-storage/internal/config"
	kvGRPC "github.com/MrTomSawyer/go-kv-storage/internal/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	cfg        *config.Config
}

func New(cfg *config.Config, handler kvGRPC.KVHandler) *App {
	gRPCServer := grpc.NewServer()
	kvGRPC.Register(gRPCServer, handler)

	return &App{
		gRPCServer: gRPCServer,
		cfg:        cfg,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.cfg.Port))
	if err != nil {
		return fmt.Errorf("%s: %w", "error initializing listener", err)
	}

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", "error serving requests", err)
	}

	log.Printf("Server started om port: %v \n", a.cfg.Port)

	return nil
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
