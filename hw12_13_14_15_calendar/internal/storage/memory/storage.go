package memorystorage

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrFailCreateEvent = errors.New("error creating event")
	ErrFailUpdateEvent = errors.New("error updating event")
	ErrNotExistEvent   = errors.New("event not exist")
)

type Event struct {
	// UUID - уникальный идентификатор события.
	UUID string
	// Заголовок - короткий текст.
	Summary string
	// Дата и время начала события.
	StartedAt string
	// Дата и время начала события.
	FinishedAt string
	// Описание события - длинный текст, опционально.
	Description string
	// UUID пользователя, владельца события.
	UserUUID string
	// Дата и время уведомления о событии.
	NotificationAt string
}

type EventUpdateAttributes struct {
	// Заголовок - короткий текст.
	Summary string
	// Дата и время начала события.
	StartedAt string
	// Дата и время начала события.
	FinishedAt string
	// Описание события - длинный текст, опционально.
	Description string
	// Get пользователя, владельца события.
	UserUUID string
	// Дата и время уведомления о событии.
	NotificationAt string
}

type EventStorage interface {
	Create(attributes Event) (bool, error)
	Update(uuid string, attributes EventUpdateAttributes) (bool, error)
	Select() (map[string]Event, error)
	Get(uuid string) (Event, error)
}

type Storage struct {
	mu     sync.RWMutex
	events map[string]Event
}

func New() EventStorage {
	return &Storage{
		mu:     sync.RWMutex{},
		events: map[string]Event{},
	}
}

func (s *Storage) Connect(ctx context.Context) error {
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	return nil
}

// Создает новое событие.
func (s *Storage) Create(e Event) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[e.UUID]; ok {
		return false, ErrFailCreateEvent
	}
	s.events[e.UUID] = e

	return true, nil
}

// Обновляет существующее в хранилище событие.
func (s *Storage) Update(uuid string, attributes EventUpdateAttributes) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	event, ok := s.events[uuid]
	if !ok {
		return false, ErrFailUpdateEvent
	}
	event.Summary = attributes.Summary
	event.StartedAt = attributes.StartedAt
	event.FinishedAt = attributes.FinishedAt
	event.Description = attributes.Description
	event.UserUUID = attributes.UserUUID
	event.NotificationAt = attributes.NotificationAt

	s.events[uuid] = event

	return true, nil
}

// Возвращает список соответствующих условию событий из хранилища, проиндексированные по идентификатору.
func (s *Storage) Select() (map[string]Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.events, nil
}

// Возвращает событие из хранилища по идентификатору.
func (s *Storage) Get(uuid string) (Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var event Event
	var ok bool

	if event, ok = s.events[uuid]; ok {
		return event, nil
	}

	return event, ErrNotExistEvent
}
