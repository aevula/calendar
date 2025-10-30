package main

import (
	"os"

	"go.uber.org/zap"

	"github.com/aevula/interview-hustlers-calendar/internal/application/config"
	logging "github.com/aevula/interview-hustlers-calendar/internal/application/log/zap"
)

func main() {
	cfg := config.MustLoad()

	logger := logging.MustLoad(cfg)
	defer logger.Sync()

	logger.Info(
		"Starting...",
		zap.Int("pid", os.Getpid()),
		zap.String("env", cfg.Env),
		zap.String("log_level", logger.Level().String()),
	)

	logger.Info("Started")
}
