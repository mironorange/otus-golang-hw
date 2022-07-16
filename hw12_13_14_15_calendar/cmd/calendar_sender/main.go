package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/broker"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/sender"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	uri := "amqp://guest:guest@localhost:5672/"

	receiver, _ := broker.New("events-broker", uri)

	if err := receiver.Connect(ctx, "events-exchange", "direct", "events"); err != nil {
		log.Fatal(err)
	}
	msgs, err := receiver.Consume(ctx, "events")
	if err != nil {
		log.Fatal(err)
	}

	sndr := sender.New(logger.New("info"))
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
