package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"
	"time"

	mqbroker "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/broker"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/pb"
	"google.golang.org/grpc"
)

var (
	configFile     string
	calendarClient pb.CalendarClient
	configuration  *config.SchedulerConfiguration
	broker         *mqbroker.Broker
)

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/scheduler.json", "Path to configuration file")
}

func main() {
	flag.Parse()

	configuration = config.NewSchedulerConfiguration()
	ctxConfig := context.TODO()
	if err := config.LoadConfig(ctxConfig, configuration, configFile); err != nil {
		log.Fatal(err)
	}

	ctx, cancelFunc := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancelFunc()

	ticker := time.NewTicker(60 * time.Second)

	// Воспользоваться gRPC соединением для того, чтобы получить события, о которых следует уведомить
	calendarAddr := net.JoinHostPort(configuration.EventsService.Host, configuration.EventsService.Port)
	log.Println(fmt.Sprintf("Connect to Calendar: %s", calendarAddr))
	grpcConnect, err := grpc.Dial(calendarAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	calendarClient = pb.NewCalendarClient(grpcConnect)

	log.Println(fmt.Sprintf("Connect to Broker: %s", configuration.Queue.URI))
	broker, _ = mqbroker.New("events-broker", configuration.Queue.URI)
	if err := broker.Connect(
		context.TODO(),
		configuration.Queue.ExchangeName,
		configuration.Queue.ExchangeType,
		configuration.Queue.QueueName,
	); err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				submitForDispatch()
				deleteOldEvents()
			}
		}
	}()

	<-ctx.Done()
	_ = grpcConnect.Close()
	_ = broker.Close(context.TODO())
}

func submitForDispatch() {
	now := time.Now().Unix()
	notifyFrom := now - now%60
	notifyTo := notifyFrom + 60

	ctx := context.TODO()
	res, err := calendarClient.GetEventsToBeNotified(
		ctx,
		&pb.GetEventsToBeNotifiedRequest{
			From: int32(notifyFrom),
			To:   int32(notifyTo),
		},
	)
	if err == nil {
		for _, e := range res.Items {
			eventMessage := mqbroker.Event{
				UUID:           e.Uuid,
				Summary:        e.Summary,
				StartedAt:      e.StartedAt,
				FinishedAt:     e.FinishedAt,
				Description:    e.Description,
				UserUUID:       e.UserUuid,
				NotificationAt: e.NotificationAt,
			}
			if body, err := eventMessage.MarshalJSON(); err == nil {
				_ = broker.Publish(context.TODO(), configuration.Queue.ExchangeName, "", body)
			}
		}
	}
}

func deleteOldEvents() {
	deleteFrom := time.Now().Add(configuration.TTL).Unix()
	ctx := context.TODO()
	res, err := calendarClient.GetOldestEvents(
		ctx,
		&pb.GetOldestEventsRequest{
			EndedAt: int32(deleteFrom),
		},
	)
	if err != nil {
		return
	}
	for _, e := range res.Items {
		_, _ = calendarClient.DeleteEvent(ctx, &pb.DeleteEventRequest{
			Uuid: e.Uuid,
		})
	}
}
