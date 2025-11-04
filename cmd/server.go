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

	app := application.New(cfg)
	app.Init(ctx)

	select {
	case <-app.Run(ctx):
	case <-ctx.Done():
	}

	ctx, stop = signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app.Shutdown(ctx)
}
