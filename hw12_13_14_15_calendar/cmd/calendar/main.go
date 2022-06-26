package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage/memory"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	// Инициализирую конфигурацию приложения
	config := NewConfig()
	config.DoSomething()

	// Инициализирую логирование приложения
	logging := logger.New(config.Logger.Level)
	logging.DoSomething()

	// Инициализирую хранилище событий в приложении
	storage := memorystorage.New()
	_, err := storage.UUID("test")
	fmt.Println(err)

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
