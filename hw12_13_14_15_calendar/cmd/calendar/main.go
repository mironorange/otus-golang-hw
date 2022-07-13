package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/grpcserver"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage/sql"
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

	// Инициализирую сервер приложения
	server := internalhttp.NewServer(net.JoinHostPort(config.Server.Host, config.Server.Port), logging, calendar)

	grpcServer := grpc.NewServer()
	grpcService := new(grpcserver.Service)
	grpcService.SetApp(calendar)

	pb.RegisterCalendarServer(grpcServer, grpcService)

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

		grpcServer.Stop()
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

		if listener, err := net.Listen("tcp", ":50051"); err == nil {
			logging.Info(fmt.Sprintf("Starting gRPC server on %s", listener.Addr().String()))
			if err := grpcServer.Serve(listener); err != nil {
				logging.Error("failed to start http server: " + err.Error())
				cancelFunc()
			}
		}
	}()

	wg.Wait()
}
