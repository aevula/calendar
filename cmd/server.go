package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/aevula/interview-hustlers-calendar/internal/application/config"
	logging "github.com/aevula/interview-hustlers-calendar/internal/application/log/zap"
	"github.com/aevula/interview-hustlers-calendar/internal/db/postgres"
)

func main() {
	cfg := config.MustLoad()

	logger := logging.MustLoad(cfg)
	defer logger.Sync()

	logger.Info(
		"Starting ...",
		zap.Int("pid", os.Getpid()),
		zap.String("env", cfg.Env),
		zap.String("log_level", logger.Level().String()),
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go watchSignals(ctx, cancel)

	initCtx, initCancel := context.WithTimeout(ctx, cfg.InitTimeout)

	db, err := postgres.New(initCtx, cfg)
	if err != nil {
		logger.Error(err.Error())
		cancel()
	}

	initCancel()

	logger.Info("Started")
	defer logger.Info("Shut down")

	go func() {
		for { // server mock
		}
	}()
	<-ctx.Done()

	logger.Info("Shutting down ...")

	shutCtx, shutCancel := context.WithTimeout(ctx, cfg.ShutTimeout)

	err = db.Close(shutCtx)
	if err != nil {
		logger.Error(err.Error())
	}

	shutCancel()
}

func watchSignals(ctx context.Context, cancel context.CancelFunc) {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-exit:
		cancel()
	case <-ctx.Done():
		return
	}
}
