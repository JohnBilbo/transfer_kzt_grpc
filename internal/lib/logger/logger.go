package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"transfer_kzt_grpc/internal/config"
	"transfer_kzt_grpc/internal/constants"
)

func NewLogger(cfg *config.Config) *zap.Logger {
	var logLevel zapcore.Level
	switch cfg.LogLevel {
	case constants.LogLevelDebug:
		logLevel = zapcore.DebugLevel
	case constants.LogLevelInfo:
		logLevel = zapcore.InfoLevel
	case constants.LogLevelWarn:
		logLevel = zapcore.WarnLevel
	case constants.LogLevelError:
		logLevel = zapcore.ErrorLevel
	default:
		logLevel = zapcore.ErrorLevel
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:      "time",
		LevelKey:     "level",
		MessageKey:   "msg",
		CallerKey:    "caller",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.CapitalColorLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	writer := zapcore.Lock(os.Stdout)
	level := logLevel
	core := zapcore.NewCore(encoder, writer, level)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger
}
