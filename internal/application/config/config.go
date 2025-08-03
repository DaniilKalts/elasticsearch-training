package config

import (
	"errors"
	"fmt"

	"github.com/goforj/godump"
	"github.com/spf13/viper"
)

type HttpConfig struct {
	Port int `mapstructure:"http_port"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"pg_host"`
	Port     int    `mapstructure:"pg_port"`
	User     string `mapstructure:"pg_user"`
	Password string `mapstructure:"pg_password"`
	Name     string `mapstructure:"pg_name"`
	SSLMode  string `mapstructure:"pg_sslmode"`
}

type Config struct {
	Server   HttpConfig     `mapstructure:"server"`
	Database PostgresConfig `mapstructure:"postgres"`
}

var cfg Config

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetDefault("http_port", 8080)

	v.SetDefault("pg_host", "localhost")
	v.SetDefault("pg_port", 5432)
	v.SetDefault("pg_user", "postgres")
	v.SetDefault("pg_password", "your_password")
	v.SetDefault("pg_name", "elastic")
	v.SetDefault("pg_sslmode", "disable")

	v.SetConfigName("app.dev")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		if !errors.Is(err, viper.ConfigFileNotFoundError{}) {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
		fmt.Println("Loaded configuration from default values")
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	if v.ConfigFileUsed() != "" {
		fmt.Printf("Loaded configuration from: %s\n", v.ConfigFileUsed())
	}

	godump.Dump(cfg)

	return &cfg, nil
}
