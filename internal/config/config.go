package config

import "time"

type Config struct {
	StoragePath string     `yaml:"storage_path"`
	GRPCConfig  GRPCConfig `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func DefaultConfig() Config {
	return Config{
		StoragePath: "./data",
		GRPCConfig: GRPCConfig{
			Port:    "50050",
			Timeout: 5 * time.Second,
		},
	}
}
