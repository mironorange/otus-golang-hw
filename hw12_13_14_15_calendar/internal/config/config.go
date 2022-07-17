package config

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

type CalendarConfiguration struct {
	Logger    LoggerConf
	Server    HTTPServerConf
	RPCServer RPCServerConf
	Events    EventsRepoConf
	Database  DatabaseConf
}

type SchedulerConfiguration struct {
	Logger        LoggerConf
	EventsService EventsServiceConf
	Queue         QueueConf
	TTL           time.Duration `config:"ttl"`
}

type SenderConfiguration struct {
	Logger LoggerConf
	Queue  QueueConf
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

type EventsServiceConf struct {
	Host    string        `config:"eventsservice-host"`
	Port    string        `config:"eventsservice-port"`
	Timeout time.Duration `config:"eventsservice-timeout"`
}

type QueueConf struct {
	Uri          string `config:"queue-uri"`
	ExchangeName string `config:"queue-exchangename"`
	ExchangeType string `config:"queue-exchangetype"`
	QueueName    string `config:"queue-queuename"`
}

func NewCalendarConfiguration() *CalendarConfiguration {
	return &CalendarConfiguration{
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

func NewSchedulerConfiguration() *SchedulerConfiguration {
	return &SchedulerConfiguration{
		EventsService: EventsServiceConf{
			Host:    "",
			Port:    "50051",
			Timeout: time.Second * 3,
		},
		Queue: QueueConf{
			Uri:          "amqp://guest:guest@localhost:5672/",
			ExchangeName: "events-exchange",
			ExchangeType: "direct",
			QueueName:    "events",
		},
		TTL: time.Hour * 24 * 365,
	}
}

func NewSenderConfiguration() *SenderConfiguration {
	return &SenderConfiguration{
		Logger: LoggerConf{
			Level: "info",
		},
		Queue: QueueConf{
			Uri:          "amqp://guest:guest@localhost:5672/",
			ExchangeName: "events-exchange",
			ExchangeType: "direct",
			QueueName:    "events",
		},
	}
}

func LoadConfig(ctx context.Context, cfg interface{}, path string) error {
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
