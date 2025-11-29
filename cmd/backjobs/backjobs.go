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

	initCtx, inited := context.WithTimeout(context.Background(), cfg.Backjobs.InitTimeout)
	app := application.NewBackjobs(initCtx, cfg)
	inited()

	app.Start(ctx)
	<-ctx.Done()

	ctx, stop = context.WithTimeout(context.Background(), cfg.Backjobs.ShutTimeout)
	defer stop()

	app.Stop(ctx)
}
