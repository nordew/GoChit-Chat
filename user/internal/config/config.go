package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net/url"
	"sync"
)

type PGConfig struct {
	PGUser     string `env:"PG_USER"`
	PGPassword string `env:"PG_PASSWORD"`
	PGHost     string `env:"PG_HOST"`
	PGDatabase string `env:"PG_DATABASE"`
	PGSSLMode  string `env:"PG_SSL_MODE"`
}

type Config struct {
	PGConfig
	GRPCPort  int    `env:"GRPC_PORT"`
	JWTSecret string `env:"JWT_SECRET"`
	Salt      string `env:"SALT"`
}

var (
	config *Config
	once   sync.Once
)

func GetConfig() (*Config, error) {
	once.Do(func() {
		config = &Config{}
		if err := cleanenv.ReadEnv(config); err != nil {
			log.Fatalf("failed to parse configs: %v", err)
		}
	})

	return config, nil
}

func MakePGConn(cfg *Config) string {
	connectionString := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=%s",
		url.QueryEscape(cfg.PGUser),
		url.QueryEscape(cfg.PGPassword),
		url.QueryEscape(cfg.PGDatabase),
		url.QueryEscape(cfg.PGSSLMode),
	)

	return connectionString
}
