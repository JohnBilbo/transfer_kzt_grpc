package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	ConfigPath string     `yaml:"config_path"`
	GRPCConfig GRPCConfig `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func DefaultConfig() *Config {
	return &Config{
		GRPCConfig: GRPCConfig{
			Port:    "50505",
			Timeout: 5 * time.Second,
		},
	}
}

func NewConfig() *Config {
	return &Config{
		GRPCConfig: GRPCConfig{
			Port:    viper.GetString("app.grpc_port"),
			Timeout: viper.GetDuration("app.grpc_timeout"),
		},
	}
}

// FindProjectRoot - функция перечисляет директории и возвращает путь, найденного корня
func FindProjectRoot(startDir string) (string, error) {
	for {
		if _, err := os.Stat(filepath.Join(startDir, "transfer_kzt_grpc")); !os.IsNotExist(err) {
			return startDir, nil
		}
		parentDir := filepath.Dir(startDir)
		if parentDir == startDir {
			return "", fmt.Errorf("project root not found")
		}
		startDir = parentDir
	}
}
