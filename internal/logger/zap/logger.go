package zap

import (
	"strings"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var toZapLevels = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

const defaultLevel = zapcore.InfoLevel

func MustLoad(cfg config.Config) *zap.Logger {
	level, ok := toZapLevels[strings.ToLower(cfg.Log.Level)]
	if !ok {
		level = defaultLevel
	}

	encoding := "json"
	if cfg.Log.Plain {
		encoding = "console"
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	if cfg.Env == config.Dev || cfg.Env == config.Test {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	outs := []string{"stdout"}
	errs := []string{"stderr"}

	initialFields := map[string]any{}

	logCfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       cfg.Env == config.Dev,
		DisableStacktrace: !cfg.Log.Trace,
		Sampling:          nil,
		Encoding:          encoding,
		EncoderConfig:     encoderCfg,
		OutputPaths:       outs,
		ErrorOutputPaths:  errs,
		InitialFields:     initialFields,
	}

	logger := zap.Must(logCfg.Build())
	return logger
}
