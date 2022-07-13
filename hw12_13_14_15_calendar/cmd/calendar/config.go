package main

import (
	"context"
	"os"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Logger    LoggerConf
	Server    HTTPServerConf
	RPCServer RPCServerConf
	Events    EventsRepoConf
	Database  DatabaseConf
}

type HTTPServerConf struct {
	Host    string        `config:"server-host"`
	Port    string        `config:"server-port"`
	Timeout time.Duration `config:"server-timeout"`
}

type RPCServerConf struct {
	Host    string        `config:"rpcserver-host"`
	Port    string        `config:"rpcserver-port"`
	Timeout time.Duration `config:"rpcserver-timeout"`
}

type EventsRepoConf struct {
	Storage string `config:"events-storage"`
}

type LoggerConf struct {
	Level string `config:"logger-level"`
}

type DatabaseConf struct {
	Driver string `config:"database-driver"`
	Dsn    string `config:"database-dsn"`
}

func NewConfig() *Config {
	return &Config{
		Server: HTTPServerConf{
			Host:    "",
			Port:    "8080",
			Timeout: time.Second * 3,
		},
		RPCServer: RPCServerConf{
			Host:    "",
			Port:    "50051",
			Timeout: time.Second * 3,
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
