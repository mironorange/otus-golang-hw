package main

import (
	"context"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
	"os"
	"time"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Logger   LoggerConf
	Server   ServerConf
	Timeout  time.Duration `config:"timeout"`
	Events   EventsRepoConf
	Database struct {
		Driver string `config:"database-driver"`
		Dsn    string `config:"database-dsn"`
	}
}

type ServerConf struct {
	Host string `config:"server-host"`
	Port string `config:"server-port"`
}

type EventsRepoConf struct {
	Storage string `config:"events-storage"`
}

type LoggerConf struct {
	Level string `config:"logger-level"`
}

func NewConfig() *Config {
	return &Config{
		Server: ServerConf{
			Host: "",
			Port: "8080",
		},
		Events: EventsRepoConf{
			Storage: "inmemory",
		},
	}
}

func LoadConfig(ctx context.Context, cfg *Config, path string) error {
	backends := make([]backend.Backend, 0)
	if _, err := os.Stat(path); err == nil {
		backends = append(backends, file.NewBackend(path))
	}
	backends = append(backends, env.NewBackend())
	backends = append(backends, flags.NewBackend())

	loader := confita.NewLoader(backends...)
	if err := loader.Load(ctx, cfg); err != nil {
		return err
	}

	return nil
}
