package memorystorage

import (
	"context"
	"errors"
	"sync"

	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
)

var (
	ErrFailCreateEvent = errors.New("error creating event")
	ErrFailUpdateEvent = errors.New("error updating event")
	ErrNotExistEvent   = errors.New("event not exist")
)

type Storage struct {
	mu     sync.RWMutex
	events map[string]storage.Event
}

func New() storage.EventStorage {
	return &Storage{
		sync.RWMutex{},
		make(map[string]storage.Event),
	}
}

func (s *Storage) Connect(ctx context.Context) error {
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	return nil
}

// Создает новое событие.
func (s *Storage) CreateEvent(
	ctx context.Context,
	uuid string,
	summary string,
	startedAt int32,
	finishedAt int32,
	description string,
	userUUID string,
	notificationAt int32,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	event := storage.Event{
		UUID:           uuid,
		Summary:        summary,
		StartedAt:      startedAt,
		FinishedAt:     finishedAt,
		Description:    description,
		UserUUID:       userUUID,
		NotificationAt: notificationAt,
	}
	if _, ok := s.events[event.UUID]; ok {
		return ErrFailCreateEvent
	}
	s.events[event.UUID] = event

	return nil
}

// Обновляет существующее в хранилище событие.
func (s *Storage) UpdateEvent(
	ctx context.Context,
	uuid string,
	summary string,
	startedAt int32,
	finishedAt int32,
	description string,
	userUUID string,
	notificationAt int32,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	event, ok := s.events[uuid]
	if !ok {
		return ErrFailUpdateEvent
	}
	event.Summary = summary
	event.StartedAt = startedAt
	event.FinishedAt = finishedAt
	event.Description = description
	event.UserUUID = userUUID
	event.NotificationAt = notificationAt

	s.events[uuid] = event

	return nil
}

// Возвращает список соответствующих условию событий из хранилища, проиндексированные по идентификатору.
func (s *Storage) GetEvents(
	ctx context.Context,
	sinceNotificationAt int32,
) ([]storage.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	events := make([]storage.Event, 0, len(s.events))
	for _, event := range s.events {
		if sinceNotificationAt < event.NotificationAt {
			events = append(events, event)
		}
	}
	return events, nil
}

// Возвращает список событий, закончившихся раньше чем наступило время endedAt.
func (s *Storage) GetOldestEvents(ctx context.Context, endedAt int32) ([]storage.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	events := make([]storage.Event, 0, len(s.events))
	for _, event := range s.events {
		if endedAt > event.FinishedAt {
			events = append(events, event)
		}
	}
	return events, nil
}

// Возвращает событие из хранилища по идентификатору.
func (s *Storage) GetEventByUUID(ctx context.Context, uuid string) (storage.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var event storage.Event
	if e, ok := s.events[uuid]; ok {
		return e, nil
	}
	return event, ErrNotExistEvent
}

func (s *Storage) DeleteEvent(ctx context.Context, uuid string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.events, uuid)
	return nil
}
