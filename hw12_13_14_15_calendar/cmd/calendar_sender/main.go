package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/broker"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/sender"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/sender.json", "Path to configuration file")
}

func main() {
	flag.Parse()

	c := config.NewSenderConfiguration()
	ctxConfig := context.TODO()
	if err := config.LoadConfig(ctxConfig, c, configFile); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	log.Println(fmt.Sprintf("Connect to Broker: %s", c.Queue.URI))
	receiver, _ := broker.New("events-broker", c.Queue.URI)

	if err := receiver.Connect(ctx, c.Queue.ExchangeName, c.Queue.ExchangeType, c.Queue.QueueName); err != nil {
		log.Fatal(err)
	}
	msgs, err := receiver.Consume(ctx, c.Queue.QueueName)
	if err != nil {
		log.Fatal(err)
	}

	sndr := sender.New(logger.New(c.Logger.Level))
	for m := range msgs {
		e := broker.Event{}
		if err := e.UnmarshalJSON(m.Data); err != nil {
			continue
		}
		_ = sndr.Send(sender.Notification{
			Title: fmt.Sprintf("Уведомление о событии: %s", e.Summary),
			Body:  e.Description,
		})
	}
}
