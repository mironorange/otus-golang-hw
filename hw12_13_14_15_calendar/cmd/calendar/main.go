package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/app"
	internalgrpc "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/grpcserver"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/server/http"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage/sql"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.json", "Path to configuration file")
}

func NewStorage(c *Config) (s storage.EventStorage) {
	switch c.Events.Storage {
	case storage.InMemoryStorageType:
		return memorystorage.New()
	case storage.SQLStorageType:
		driver := c.Database.Driver
		dsn := c.Database.Dsn
		s := sqlstorage.New(driver, dsn)
		ctx := context.TODO()
		_ = s.Connect(ctx)
		return s.(storage.EventStorage)
	default:
		log.Fatal("storage not configured")
	}
	return s
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

	// Инициализирую объект приложения
	calendar := app.New(logging, NewStorage(config))

	// Инициализирую сервер приложения
	server := internalhttp.NewServer(net.JoinHostPort(config.Server.Host, config.Server.Port), logging, calendar)
	grpcServer := internalgrpc.NewServer(net.JoinHostPort(config.RPCServer.Host, config.RPCServer.Port), logging, calendar)

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
		if err := grpcServer.Stop(ctx); err != nil {
			logging.Error("failed to stop grpc server: " + err.Error())
		}
	}()

	logging.Info("calendar is running...")

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		logging.Info("HTTP Server is running...")
		if err := server.Start(ctx); err != nil {
			logging.Error("failed to start http server: " + err.Error())
			cancelFunc()
		}
	}()

	go func() {
		defer wg.Done()
		logging.Info("gRPC Server is running...")
		if err := grpcServer.Start(ctx); err != nil {
			logging.Error("failed to start grpc server: " + err.Error())
			cancelFunc()
		}
	}()

	wg.Wait()
}
