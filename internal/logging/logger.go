package logging

import (
	"strings"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogLevel = zapcore.Level
type LogTag = zap.Field

type Logger interface {
	Debug(msg string, fields ...LogTag)
	Info(msg string, fields ...LogTag)
	Warn(msg string, fields ...LogTag)
	Error(msg string, fields ...LogTag)
	DPanic(msg string, fields ...LogTag)
	Panic(msg string, fields ...LogTag)
	Fatal(msg string, fields ...LogTag)

	Sync() error
	Level() LogLevel
	With(fields ...LogTag) Logger

	Int(key string, val int) LogTag
	String(key string, val string) LogTag
	Bool(key string, val bool) LogTag
}

type logger struct {
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

func MustLoad(cfg config.Config) Logger {
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

	zapLogger := zap.Must(logCfg.Build())
	return &logger{Logger: zapLogger}
}

func (l *logger) With(fields ...LogTag) Logger {
	return &logger{l.Logger.With(fields...)}
}

func (l *logger) Int(key string, val int) LogTag {
	return zap.Int(key, val)
}

func (l *logger) String(key string, val string) LogTag {
	return zap.String(key, val)
}

func (l *logger) Bool(key string, val bool) LogTag {
	return zap.Bool(key, val)
}
