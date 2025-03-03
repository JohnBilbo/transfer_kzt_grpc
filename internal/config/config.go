package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	ConfigPath string
	LogLevel   string
	GRPCConfig GRPCConfig
	DBConfig   DBConfig
}

type GRPCConfig struct {
	Port    int
	Timeout time.Duration
}

type DBConfig struct {
	Host                  string
	Port                  string
	Name                  string
	User                  string
	Password              string
	SSLMode               string
	TimeZone              string
	MaxConns              int32
	MaxConnLifeTime       time.Duration
	MaxConnLifeTimeJitter time.Duration
	MaxConnIdleTime       time.Duration
}

func DefaultConfig() *Config {
	return &Config{
		GRPCConfig: GRPCConfig{
			Port:    50505,
			Timeout: 5 * time.Second,
		},
	}
}

func NewConfig() *Config {
	viper.SetConfigName("./../config/local.yaml")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	return &Config{
		LogLevel: viper.GetString("app.log_level"),
		GRPCConfig: GRPCConfig{
			Port:    viper.GetInt("app.grpc.port"),
			Timeout: viper.GetDuration("app.grpc.timeout_second"),
		},
		DBConfig: DBConfig{
			Host:                  viper.GetString("app.db.host"),
			User:                  viper.GetString("app.db.user"),
			Password:              viper.GetString("app.db.pass"),
			Name:                  viper.GetString("app.db.name"),
			Port:                  viper.GetString("app.db.port"),
			SSLMode:               viper.GetString("app.db.ssl_mode"),
			TimeZone:              viper.GetString("app.db.time_zone"),
			MaxConns:              viper.GetInt32("app.db.max_conns"),
			MaxConnLifeTime:       viper.GetDuration("app.db.max_conns_lifetime_minute"),
			MaxConnLifeTimeJitter: viper.GetDuration("app.db.max_conns_lifetime_jitter_minute"),
			MaxConnIdleTime:       viper.GetDuration("app.db.max_conns_idle_time_minute"),
		},
	}
}
