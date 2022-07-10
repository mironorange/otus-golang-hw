package main

import (
	"context"
	"fmt"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
	"log"
	"os"
	"time"
)

type Config struct {
	Port        uint32        `config:"port"`
	Timeout     time.Duration `config:"timeout"`
	Events struct{
		Storage string `config:"events-storage"`
	}
	Database struct {
		Driver string `config:"database-driver"`
		Dsn string `config:"database-dsn"`
	}
}

func main() {
	cfg := Config{}

	backends := []backend.Backend{
		env.NewBackend(),
		flags.NewBackend(),
	}

	path := "/Users/ivanmironov/go/src/github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/configs/config.json"
	if _, err := os.Stat(path); err == nil {
		backends = append(backends, file.NewBackend(path))
	}

	loader := confita.NewLoader(backends...)
	if err := loader.Load(context.Background(), &cfg); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port: ", cfg.Port)
	fmt.Println("EventsStorage.Driver: ", cfg.Events.Storage)
	fmt.Println("Database.Driver: ", cfg.Database.Driver)
	fmt.Println("Database.Dsn: ", cfg.Database.Dsn)
	//fmt.Println("Timeout: ", cfg.Timeout)
}
