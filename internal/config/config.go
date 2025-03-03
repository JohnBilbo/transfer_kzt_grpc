package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
	"time"
)

type Config struct {
	LoggerLevel zapcore.Level
	ConfigPath  string
	GRPCConfig  GRPCConfig
}

type GRPCConfig struct {
	Port    int
	Timeout time.Duration
}

func DefaultConfig() *Config {
	return &Config{
		LoggerLevel: zapcore.DebugLevel,
		GRPCConfig: GRPCConfig{
			Port:    50505,
			Timeout: 5 * time.Second,
		},
	}
}

func NewConfig() *Config {
	return &Config{
		GRPCConfig: GRPCConfig{
			Port:    viper.GetInt("app.grpc_port"),
			Timeout: viper.GetDuration("app.grpc_timeout"),
		},
	}
}
