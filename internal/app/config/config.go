package config

import (
	"errors"

	"github.com/spf13/viper"
)

type (
	// Config ...
	Config struct {
		Server   ServerConfig
		Postgres PostgresConfig
		Auth     AuthConfig
		Session  SessionConfig
	}

	// ServerConfig ...
	ServerConfig struct {
		BindAddr string
	}

	// PostgresConfig ...
	PostgresConfig struct {
		PostgresHost    string
		PostgresDbName  string
		PostgresSslMode string
	}

	// AuthConfig ...
	AuthConfig struct {
		SigninKey       string
		AccessTokenTTL  int
		RefreshTokenTTL int
	}

	// SessionConfig ...
	SessionConfig struct {
		SessionKey string
	}
)

// LoadConfig ...
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("error: config file not found")
		}
		return nil, err
	}

	return v, nil
}

// ParseConfig ...
func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
