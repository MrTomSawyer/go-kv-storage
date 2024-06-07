package main

import (
	"context"
	"github.com/MrTomSawyer/go-kv-storage/internal/app"
	"github.com/MrTomSawyer/go-kv-storage/internal/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.New()
	cfg.MustInit()

	ctx, cancel := context.WithCancel(context.Background())
	server := app.New(ctx, cfg)

	go server.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	cancel()
	server.GRPCServer.Stop()
	log.Println("server stopped")
}
