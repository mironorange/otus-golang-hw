package main

import (
	"context"
	"log"

	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/broker"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/pb"
	"google.golang.org/grpc"
)

func main() {
	// Воспользоваться gRPC соединением для того, чтобы получить события, о которых следует уведомить
	grpcConnect, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConnect.Close()
	client := pb.NewCalendarClient(grpcConnect)
	ctx := context.TODO()
	res, _ := client.GetEvents(
		ctx,
		&pb.GetEventsRequest{
			SinceNotificationAt: -1,
		},
	)

	//
	uri := "amqp://guest:guest@localhost:5672/"
	b, _ := broker.New("events-broker", uri)
	if err := b.Connect(ctx, "events-exchange", "direct", "events"); err != nil {
		log.Fatal(err)
	}
	defer b.Close(context.TODO())

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
		if eventBody, err := eventMessage.MarshalJSON(); err == nil {
			_ = b.Publish(context.TODO(), "events-exchange", "", eventBody)
		}
	}
}
