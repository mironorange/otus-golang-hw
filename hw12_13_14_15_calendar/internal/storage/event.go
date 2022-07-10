package storage

import "context"

type EventStorage interface {
	CreateEvent(
		ctx context.Context,
		uuid string,
		summary string,
		startedAt int,
		finishedAt int,
		description string,
		userUUID string,
		notificationAt int,
	) error
	UpdateEvent(
		ctx context.Context,
		uuid string,
		summary string,
		startedAt int,
		finishedAt int,
		description string,
		userUUID string,
		notificationAt int,
	) error
	GetEvents(
		ctx context.Context,
	) ([]Event, error)
	GetEventByUUID(
		ctx context.Context,
		uuid string,
	) (Event, error)
}

type Event struct {
	// UUID - уникальный идентификатор события.
	UUID string `db:"uuid"`
	// Заголовок - короткий текст.
	Summary string `db:"summary"`
	// Unix timestamp даты и времени начала события.
	StartedAt int `db:"started_at"`
	// Unix timestamp даты и времени завершения события.
	FinishedAt int `db:"finished_at"`
	// Описание события - длинный текст, опционально.
	Description string `db:"description"`
	// UUID пользователя, владельца события.
	UserUUID string `db:"user_uuid"`
	// Unix timestamp даты и времени уведомления о событии.
	NotificationAt int `db:"notification_at"`
}

type EventUpdateAttributes struct {
	// Заголовок - короткий текст.
	Summary string `db:"summary"`
	// Unix timestamp даты и времени начала события.
	StartedAt int `db:"started_at"`
	// Unix timestamp даты и времени завершения события.
	FinishedAt int `db:"finished_at"`
	// Описание события.
	Description string `db:"description"`
	// Get пользователя, владельца события.
	UserUUID string `db:"user_uuid"`
	// Unix timestamp даты и времени уведомления о событии.
	NotificationAt int `db:"notification_at"`
}
