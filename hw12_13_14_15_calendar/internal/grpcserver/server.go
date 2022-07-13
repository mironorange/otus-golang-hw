package grpcserver

import (
	"context"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	pb.UnimplementedCalendarServer
	application app.Storage
}

func (s *Service) SetApp(a app.Storage) {
	s.application = a
}

func (s *Service) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	err := s.application.CreateEvent(
		ctx,
		req.Uuid,
		req.Summary,
		int(req.StartedAt),
		int(req.FinishedAt),
		req.Description,
		req.UserUuid,
		int(req.NotificationAt),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error when trying to create event")
	}
	return &pb.CreateEventResponse{}, nil
}

func (s *Service) UpdateEvent(ctx context.Context, req *pb.UpdateEventRequest) (*pb.UpdateEventResponse, error) {
	uuid := req.Uuid
	attrs := req.Attributes
	err := s.application.UpdateEvent(
		ctx,
		uuid,
		attrs.Summary,
		int(attrs.StartedAt),
		int(attrs.FinishedAt),
		attrs.Description,
		attrs.UserUuid,
		int(attrs.NotificationAt),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error when trying to update event")
	}
	return &pb.UpdateEventResponse{}, nil
}

func (s *Service) GetEvents(ctx context.Context, req *pb.GetEventsRequest) (*pb.GetEventsResponse, error) {
	events, err := s.application.GetEvents(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error when trying to get events")
	}

	es := make([]*pb.Event, 0, len(events))
	for _, event := range events {
		es = append(es, &pb.Event{
			Uuid:           event.UUID,
			Summary:        event.Summary,
			StartedAt:      int64(event.StartedAt),
			FinishedAt:     int64(event.FinishedAt),
			Description:    event.Description,
			UserUuid:       event.UserUUID,
			NotificationAt: 0,
		})
	}

	return &pb.GetEventsResponse{Items: es}, nil
}

func (s *Service) GetEvent(ctx context.Context, req *pb.GetEventRequest) (*pb.GetEventResponse, error) {
	event, err := s.application.GetEventByUUID(ctx, req.Uuid)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Error when trying to get event")
	}

	e := &pb.Event{
		Uuid:           event.UUID,
		Summary:        event.Summary,
		StartedAt:      int64(event.StartedAt),
		FinishedAt:     int64(event.FinishedAt),
		Description:    event.Description,
		UserUuid:       event.UserUUID,
		NotificationAt: 0,
	}

	return &pb.GetEventResponse{Item: e}, nil
}
