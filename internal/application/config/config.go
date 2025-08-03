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
	Host     string `mapstructure:"db_host"`
	Port     int    `mapstructure:"db_port"`
	User     string `mapstructure:"db_user"`
	Password string `mapstructure:"db_password"`
	Name     string `mapstructure:"db_name"`
	SSLMode  string `mapstructure:"db_sslmode"`
}

type Config struct {
	Server   HttpConfig     `mapstructure:"server"`
	Database PostgresConfig `mapstructure:"database"`
}

var cfg Config

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetDefault("http_port", 8080)

	v.SetDefault("db_host", "localhost")
	v.SetDefault("db_port", 5432)
	v.SetDefault("db_user", "postgres")
	v.SetDefault("db_password", "dbpassword")
	v.SetDefault("db_name", "keycloak")
	v.SetDefault("db_sslmode", "disable")

	v.SetConfigName("app.dev")
	v.SetConfigType("yaml")
	v.AddConfigPath("/config")

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
