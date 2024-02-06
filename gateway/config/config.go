package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	HTTPPort int `env:"HTTP_PORT"`
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
