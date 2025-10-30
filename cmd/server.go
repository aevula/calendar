package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/aevula/interview-hustlers-calendar/internal/application"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := application.New(ctx)

	go app.Run(ctx)
	<-ctx.Done()

	ctx, close := context.WithTimeout(context.Background(), app.Cfg.ShutTimeout)
	defer close()

	go app.Shutdown(ctx)
	<-ctx.Done()
}
