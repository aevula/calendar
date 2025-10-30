package main

import (
	"context"
	"os"

	"go.uber.org/zap"

	"github.com/aevula/interview-hustlers-calendar/internal/application/config"
	logging "github.com/aevula/interview-hustlers-calendar/internal/application/log/zap"
	"github.com/aevula/interview-hustlers-calendar/internal/db/postgres"
)

func main() {
	cfg := config.MustLoad()

	logger := logging.MustLoad(cfg)
	defer logger.Sync()

	logger.Sugar().Info(
		"Starting...",
		zap.Int("pid", os.Getpid()),
		zap.String("env", cfg.Env),
		zap.String("log_level", logger.Level().String()),
	)

	ctx := context.Background()

	db, err := postgres.New(ctx, cfg)
	if err != nil {
		logger.Fatal(err.Error())
	}
	db.Ping(ctx)

	logger.Info("Started")
}
