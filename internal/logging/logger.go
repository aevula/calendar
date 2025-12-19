package logging

import (
	"strings"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

var toZapLevels = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

const DefaultLevel = zapcore.InfoLevel

func MustLoad(cfg config.Config) *Logger {
	level, ok := toZapLevels[strings.ToLower(cfg.Log.Level)]
	if !ok {
		level = DefaultLevel
	}

	encoding := "json"
	if cfg.Log.Plain {
		encoding = "console"
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	if cfg.IsDev() || cfg.IsTest() {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	outs := []string{"stdout"}
	errs := []string{"stderr"}

	initialFields := map[string]any{}

	logCfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       cfg.IsDev(),
		DisableStacktrace: !cfg.Log.Trace,
		Sampling:          nil,
		Encoding:          encoding,
		EncoderConfig:     encoderCfg,
		OutputPaths:       outs,
		ErrorOutputPaths:  errs,
		InitialFields:     initialFields,
	}

	logger := zap.Must(logCfg.Build())
	return &Logger{Logger: logger}
}

func (l *Logger) Int(key string, val int) zap.Field {
	return zap.Int(key, val)
}

func (l *Logger) String(key string, val string) zap.Field {
	return zap.String(key, val)
}
