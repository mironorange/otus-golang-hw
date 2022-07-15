package internalhttp

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
)

var eventsRegexp = regexp.MustCompile(`\/events\/(\w+\-\w+\-\w+\-\w+\-\w+)`)

type WrapServer struct {
	server http.Server
}

type StatusRecorder struct {
	http.ResponseWriter
	StatusCode int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.StatusCode = status
	r.ResponseWriter.WriteHeader(status)
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello, World!"))
}

func createEventsHandler(l Logger, a Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/events/" {
			// GET /events/a6e592bc-8627-4e13-b4a6-d7072864602a
			// Возвращает информацию о существующем в хранилище событии.
			if r.Method == "GET" {
				if eventsRegexp.MatchString(r.URL.Path) {
					submatch := eventsRegexp.FindStringSubmatch(r.URL.Path)
					event, err := a.GetEventByUUID(
						context.Background(),
						submatch[1],
					)
					if err == nil {
						e := Event{
							UUID:           event.UUID,
							Summary:        event.Summary,
							StartedAt:      event.StartedAt,
							FinishedAt:     event.FinishedAt,
							Description:    event.Description,
							UserUUID:       event.UserUUID,
							NotificationAt: event.NotificationAt,
						}
						jsonEvent, _ := e.MarshalJSON()
						w.Write(jsonEvent)
						return
					}
				}
			} else if r.Method == "PUT" {
				// PUT /events/a6e592bc-8627-4e13-b4a6-d7072864602a.
				// Полностью обновляет существующее в хранилище событие.
				if eventsRegexp.MatchString(r.URL.Path) {
					submatch := eventsRegexp.FindStringSubmatch(r.URL.Path)
					event, err := a.GetEventByUUID(
						context.Background(),
						submatch[1],
					)
					if err != nil {
						l.Error(fmt.Sprint(err))
						http.NotFound(w, r)
						return
					}
					attrs := EventUpdateAttributes{}
					body, _ := io.ReadAll(r.Body)
					if err := attrs.UnmarshalJSON(body); err != nil {
						w.WriteHeader(500)
						l.Error(fmt.Sprint(err))
					}
					if err := r.Body.Close(); err != nil {
						w.WriteHeader(500)
						l.Error(fmt.Sprint(err))
					}
					err = a.UpdateEvent(
						context.Background(),
						event.UUID,
						attrs.Summary,
						attrs.StartedAt,
						attrs.FinishedAt,
						attrs.Description,
						attrs.UserUUID,
						attrs.NotificationAt,
					)
					if err != nil {
						w.WriteHeader(500)
						l.Error(fmt.Sprint(err))
					}
					return
				}
			}

			http.NotFound(w, r)
			return
		}
		if r.Method == "GET" {
			events, err := a.GetEvents(context.TODO())
			if err != nil {
				w.WriteHeader(500)
				l.Error(fmt.Sprint(err))
				return
			}
			items := make(ListOfEvents, 0, len(events))
			for _, e := range events {
				items = append(items, Event{
					UUID:           e.UUID,
					Summary:        e.Summary,
					StartedAt:      e.StartedAt,
					FinishedAt:     e.FinishedAt,
					Description:    e.Description,
					UserUUID:       e.UserUUID,
					NotificationAt: e.NotificationAt,
				})
			}
			jsonEvents, _ := items.MarshalJSON()
			w.Write(jsonEvents)
			return
		} else if r.Method == "POST" {
			// POST /events/ Создает событие в базе данных.
			e := Event{}
			body, _ := io.ReadAll(r.Body)
			if err := e.UnmarshalJSON(body); err != nil {
				l.Error(fmt.Sprint(err))
			}
			if err := r.Body.Close(); err != nil {
				l.Error(fmt.Sprint(err))
			}
			err := a.CreateEvent(
				context.Background(),
				e.UUID,
				e.Summary,
				e.StartedAt,
				e.FinishedAt,
				e.Description,
				e.UserUUID,
				e.NotificationAt,
			)
			if err != nil {
				w.WriteHeader(500)
				l.Error(fmt.Sprint(err))
			}
			w.WriteHeader(201)
			return
		}
		http.NotFound(w, r)
	}
}

type Logger interface {
	Info(msg string)
	Error(msg string)
}

type Application interface {
	CreateEvent(
		ctx context.Context,
		uuid string,
		summary string,
		startedAt int32,
		finishedAt int32,
		description string,
		userUUID string,
		notificationAt int32,
	) error
	UpdateEvent(
		ctx context.Context,
		uuid string,
		summary string,
		startedAt int32,
		finishedAt int32,
		description string,
		userUUID string,
		notificationAt int32,
	) error
	GetEvents(
		ctx context.Context,
	) ([]storage.Event, error)
	GetEventByUUID(
		ctx context.Context,
		uuid string,
	) (storage.Event, error)
}

func NewServer(addr string, logger Logger, app Application) *WrapServer {
	mux := http.NewServeMux()

	helloHandler := http.HandlerFunc(hello)
	eventsHandler := createEventsHandler(logger, app)
	mux.Handle("/events/", loggingMiddleware(eventsHandler, logger))
	mux.Handle("/", loggingMiddleware(helloHandler, logger))

	return &WrapServer{
		server: http.Server{
			Addr:         addr,
			Handler:      mux,
			ReadTimeout:  3 * time.Second,
			WriteTimeout: 3 * time.Second,
		},
	}
}

func (s *WrapServer) Start(ctx context.Context) error {
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}
	<-ctx.Done()
	return nil
}

func (s *WrapServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
