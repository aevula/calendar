package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
	"github.com/aevula/interview-hustlers-calendar/internal/databases/postgres"
	logging "github.com/aevula/interview-hustlers-calendar/internal/logger/zap"
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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	initCtx, initCancel := context.WithTimeout(ctx, cfg.InitTimeout)
	defer initCancel()

	db, err := postgres.New(initCtx, cfg)
	if err != nil {
		stop()
		logger.Error(err.Error())
		return
	}

	logger.Info("Started")

	go func() {
		for { // server mock
		}
	}()
	<-ctx.Done()

	logger.Info("Shutting down ...")

	shutCtx, shutCancel := context.WithTimeout(context.Background(), cfg.ShutTimeout)
	defer shutCancel()

	err = db.Close(shutCtx)
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Shut down")
}
