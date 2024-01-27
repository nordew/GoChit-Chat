package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net/url"
)

type Config struct {
	PGUser     string `env:"PG_USER"`
	PGPassword string `env:"PG_PASSWORD"`
	PGHost     string `env:"PG_HOST"`
	PGDatabase string `env:"PG_DATABASE"`
	PGSSLMode  string `env:"PG_SSL_MODE"`

	GRPCPort int `env:"GRPC_PORT"`
}

var (
	config *Config
)

func GetConfig() (*Config, error) {
	config = &Config{}
	if err := cleanenv.ReadEnv(config); err != nil {
		log.Fatalf("failed to parse configs: %v", err)
	}

	return config, nil
}

func MakePGConn(cfg *Config) string {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		url.QueryEscape(cfg.PGUser),
		url.QueryEscape(cfg.PGPassword),
		url.QueryEscape(cfg.PGHost),
		url.QueryEscape(cfg.PGDatabase),
		url.QueryEscape(cfg.PGSSLMode),
	)

	return connectionString
}
