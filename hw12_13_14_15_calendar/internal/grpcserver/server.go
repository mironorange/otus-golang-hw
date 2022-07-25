package grpcserver

import (
	"context"
	"fmt"
	"net"

	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/pb"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WrapServer struct {
	server   *grpc.Server
	listener net.Listener
	logger   app.Logger
}

func NewServer(addr string, logger app.Logger, app app.EventStorage) *WrapServer {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Info(fmt.Sprintf("%s", err))
	}

	s := &WrapServer{
		server:   grpc.NewServer(),
		listener: listener,
		logger:   logger,
	}

	eventsService := new(EventsService)
	eventsService.SetApp(app)
	eventsService.SetLogger(logger)
	pb.RegisterCalendarServer(s.server, eventsService)

	return s
}

type EventsService struct {
	pb.UnimplementedCalendarServer
	application app.EventStorage
	logger      app.Logger
}

func (s *EventsService) SetApp(a app.EventStorage) {
	s.application = a
}

func (s *EventsService) SetLogger(l app.Logger) {
	s.logger = l
}

func (s *EventsService) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	s.logger.Info("Called CreateEvent")
	err := s.application.CreateEvent(
		ctx,
		req.Uuid,
		req.Summary,
		req.StartedAt,
		req.FinishedAt,
		req.Description,
		req.UserUuid,
		req.NotificationAt,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error when trying to create event")
	}
	return &pb.CreateEventResponse{}, nil
}

func (s *EventsService) UpdateEvent(ctx context.Context, req *pb.UpdateEventRequest) (*pb.UpdateEventResponse, error) {
	s.logger.Info("Called UpdateEvent")
	uuid := req.Uuid
	attrs := req.Attributes
	err := s.application.UpdateEvent(
		ctx,
		uuid,
		attrs.Summary,
		attrs.StartedAt,
		attrs.FinishedAt,
		attrs.Description,
		attrs.UserUuid,
		attrs.NotificationAt,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error when trying to update event")
	}
	return &pb.UpdateEventResponse{}, nil
}

func (s *EventsService) GetEvents(ctx context.Context, req *pb.GetEventsRequest) (*pb.GetEventsResponse, error) {
	s.logger.Info("Called GetEvents")
	events, err := s.application.GetEvents(ctx, req.SinceNotificationAt)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error when trying to get events")
	}
	return createGetEventsResponse(events), nil
}

func (s *EventsService) GetOldestEvents(
	ctx context.Context,
	req *pb.GetOldestEventsRequest,
) (*pb.GetEventsResponse, error) {
	s.logger.Info("Called GetOldestEvents")
	events, err := s.application.GetOldestEvents(ctx, req.EndedAt)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error when trying to get oldest events")
	}
	return createGetEventsResponse(events), nil
}

func (s *EventsService) GetEventsToBeNotified(
	ctx context.Context,
	req *pb.GetEventsToBeNotifiedRequest,
) (*pb.GetEventsResponse, error) {
	s.logger.Info("Called GetEventsToBeNotified")
	events, err := s.application.GetEventsToBeNotified(ctx, req.From, req.To)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error when trying to get events to be notified")
	}
	return createGetEventsResponse(events), nil
}

func createGetEventsResponse(events []storage.Event) *pb.GetEventsResponse {
	items := make([]*pb.Event, 0, len(events))
	for _, event := range events {
		items = append(items, &pb.Event{
			Uuid:           event.UUID,
			Summary:        event.Summary,
			StartedAt:      event.StartedAt,
			FinishedAt:     event.FinishedAt,
			Description:    event.Description,
			UserUuid:       event.UserUUID,
			NotificationAt: event.NotificationAt,
		})
	}
	return &pb.GetEventsResponse{Items: items}
}

func (s *EventsService) GetEvent(ctx context.Context, req *pb.GetEventRequest) (*pb.GetEventResponse, error) {
	s.logger.Info("Called GetEvent")
	event, err := s.application.GetEventByUUID(ctx, req.Uuid)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Error when trying to get event")
	}

	e := &pb.Event{
		Uuid:           event.UUID,
		Summary:        event.Summary,
		StartedAt:      event.StartedAt,
		FinishedAt:     event.FinishedAt,
		Description:    event.Description,
		UserUuid:       event.UserUUID,
		NotificationAt: event.NotificationAt,
	}

	return &pb.GetEventResponse{Item: e}, nil
}

func (s *EventsService) DeleteEvent(ctx context.Context, req *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {
	s.logger.Info("Called DeleteEvent")
	err := s.application.DeleteEvent(ctx, req.Uuid)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error when trying to delete event")
	}

	return &pb.DeleteEventResponse{}, nil
}

func (s *WrapServer) Start(ctx context.Context) error {
	s.logger.Info(fmt.Sprintf("listen rpc server: %s", s.listener.Addr()))
	err := s.server.Serve(s.listener)
	<-ctx.Done()
	return err
}

func (s *WrapServer) Stop(ctx context.Context) error {
	s.server.GracefulStop()
	return nil
}
