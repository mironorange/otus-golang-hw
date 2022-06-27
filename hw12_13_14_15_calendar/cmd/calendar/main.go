package main

import (
	"context"
	"flag"
	memorystorage "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/server/http"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.json", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	// Инициализирую конфигурацию приложения
	config := NewConfig()
	ctxConfig := context.TODO()
	if err := LoadConfig(ctxConfig, config, configFile); err != nil {
		log.Fatal(err)
	}

	// Инициализирую логирование приложения
	logging := logger.New(config.Logger.Level)
	logging.DoSomething()

	var storage app.Storage
	switch config.Events.Storage {
	case "inmemory":
		s := memorystorage.New()
		storage = s.(app.Storage)
	case "database":
		driver := config.Database.Driver
		dsn := config.Database.Dsn
		s := sqlstorage.New(driver, dsn)
		ctxStorage := context.TODO()
		if err := s.Connect(ctxStorage); err != nil {
			log.Fatal(err)
		}
		storage = s.(app.Storage)
	default:
		log.Fatal("storage not configured")
	}

	// Инициализирую объект приложения
	calendar := app.New(logging, storage)
	calendar.DoSomething()

	// Инициализирую сервер приложения
	server := internalhttp.NewServer(logging, calendar)
	server.DoSomething()
	ctx, cancelFunc := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancelFunc()

	// Обрабатываю запросы к серверу приложения
	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logging.Error("failed to stop http server: " + err.Error())
		}
	}()

	logging.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logging.Error("failed to start http server: " + err.Error())
		cancelFunc()
		os.Exit(1) //nolint:gocritic
	}
}
