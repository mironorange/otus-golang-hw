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

	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/broker"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/pb"
	"google.golang.org/grpc"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/scheduler.json", "Path to configuration file")
}

func main() {
	flag.Parse()

	c := config.NewSchedulerConfiguration()
	ctxConfig := context.TODO()
	if err := config.LoadConfig(ctxConfig, c, configFile); err != nil {
		log.Fatal(err)
	}

	ctx, cancelFunc := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancelFunc()

	ticker := time.NewTicker(5 * time.Second)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				repeatRun(c)
			}
		}
	}()

	<-ctx.Done()
}

func repeatRun(c *config.SchedulerConfiguration) {
	// Воспользоваться gRPC соединением для того, чтобы получить события, о которых следует уведомить
	grpcConnect, err := grpc.Dial(net.JoinHostPort(c.EventsService.Host, c.EventsService.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConnect.Close()
	client := pb.NewCalendarClient(grpcConnect)

	b, _ := broker.New("events-broker", c.Queue.URI)
	if err := b.Connect(context.TODO(), c.Queue.ExchangeName, c.Queue.ExchangeType, c.Queue.QueueName); err != nil {
		log.Fatal(err)
	}
	defer b.Close(context.TODO())

	ctx := context.TODO()
	res, err := client.GetEvents(
		ctx,
		&pb.GetEventsRequest{
			SinceNotificationAt: -1,
		},
	)
	if err == nil {
		for _, e := range res.Items {
			eventMessage := broker.Event{
				UUID:           e.Uuid,
				Summary:        e.Summary,
				StartedAt:      e.StartedAt,
				FinishedAt:     e.FinishedAt,
				Description:    e.Description,
				UserUUID:       e.UserUuid,
				NotificationAt: e.NotificationAt,
			}
			if body, err := eventMessage.MarshalJSON(); err == nil {
				_ = b.Publish(context.TODO(), c.Queue.ExchangeName, "", body)
			}
		}
	}

	now := time.Now().Add(c.TTL).Unix()
	ctx = context.TODO()
	res, err = client.GetOldestEvents(
		ctx,
		&pb.GetOldestEventsRequest{
			EndedAt: int32(now),
		},
	)
	if err == nil {
		for _, e := range res.Items {
			fmt.Println(res)
			_, err = client.DeleteEvent(ctx, &pb.DeleteEventRequest{
				Uuid: e.Uuid,
			})
			fmt.Println(err)
		}
	}
}
