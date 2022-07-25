package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/config"
	internalgrpc "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/grpcserver"
	logs "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/server/http"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage/sql"
)

var (
	configFile string
	logger     *logs.WrapLogger
)

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
	c := config.NewCalendarConfiguration()
	ctxConfig := context.TODO()
	if err := config.LoadConfig(ctxConfig, c, configFile); err != nil {
		log.Fatal(err)
	}

	// Инициализирую логирование приложения
	logger = logs.New(c.Logger.Level)

	// Инициализирую объект приложения
	calendar := app.New(logger, NewStorage(c))

	// Инициализирую сервер приложения
	server := internalhttp.NewServer(net.JoinHostPort(c.Server.Host, c.Server.Port), logger, calendar)
	grpcServer := internalgrpc.NewServer(net.JoinHostPort(c.RPCServer.Host, c.RPCServer.Port), logger, calendar)

	ctx, cancelFunc := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancelFunc()

	// Обрабатываю запросы к серверу приложения
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logger.Error("failed to stop http server: " + err.Error())
		}
		if err := grpcServer.Stop(ctx); err != nil {
			logger.Error("failed to stop grpc server: " + err.Error())
		}
	}()

	logger.Info("calendar is running...")

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := server.Start(ctx); err != nil {
			logger.Error("failed to start http server: " + err.Error())
			cancelFunc()
		}
	}()

	go func() {
		defer wg.Done()
		if err := grpcServer.Start(ctx); err != nil {
			logger.Error("failed to start grpc server: " + err.Error())
			cancelFunc()
		}
	}()

	wg.Wait()
}

func NewStorage(c *config.CalendarConfiguration) (s storage.EventStorage) {
	switch c.Events.Storage {
	case storage.InMemoryStorageType:
		logger.Info("Connect to inmemory")
		return memorystorage.New()
	case storage.SQLStorageType:
		driver := c.Database.Driver
		dsn := c.Database.Dsn
		logger.Info(fmt.Sprintf("Connect to %s", dsn))
		s := sqlstorage.New(driver, dsn)
		ctx := context.Background()
		if err := s.Connect(ctx); err != nil {
			logger.Error(fmt.Sprintf("%s", err))
		}
		return s.(storage.EventStorage)
	default:
		panic("storage not configured")
	}
}
