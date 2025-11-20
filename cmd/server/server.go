package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/aevula/interview-hustlers-calendar/internal/application"
	"github.com/aevula/interview-hustlers-calendar/internal/config"
)

func main() {
	cfg := config.MustLoad()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	ctx, stop = context.WithTimeout(ctx, cfg.Server.InitTimeout)
	defer stop()

	app := application.NewServer(cfg)
	app.Init(ctx)

	select {
	case <-app.Run(ctx):
	case <-ctx.Done():
	}

	ctx, stop = context.WithTimeout(context.Background(), cfg.Server.ShutTimeout)
	defer stop()

	app.Stop(ctx)
}
