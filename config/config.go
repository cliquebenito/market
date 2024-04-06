package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type DataBaseconfig struct {
	User        string `yaml:"POSTGRES_USER"`
	Password    string `yaml:"POSTGRES_PASSWORD"`
	Name        string `yaml:"POSTGRES_DB"`
	Host        string `yaml:"POSTGRES_HOST"`
	Port        string `yaml:"POSTGRES_PORT"`
	SSLMode     string `yaml:"ssl_mode"`
	MaxAttempts int    `yaml:"max_attempts"`
	PG_URL      string `yaml:"pg_url"`
}

type ServeConfig struct {
	Port string `yaml:"port"`
}

type Config struct {
	Database DataBaseconfig `yaml:"database"`
	Api      ServeConfig    `yaml:"api"`
}

// // отдаёт нам конфиг исходя из config.yaml
func CheckConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("../config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
