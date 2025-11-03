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

	app.Run(ctx)
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), app.Cfg.ShutTimeout)
	defer cancel()

	app.Shutdown(ctx)
}
