package storage

import "context"

type EventStorage interface {
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
		sinceNotificationAt int32,
	) ([]Event, error)
	GetOldestEvents(
		ctx context.Context,
		endedAt int32,
	) ([]Event, error)
	GetEventsToBeNotified(
		ctx context.Context,
		from int32,
		to int32,
	) ([]Event, error)
	GetEventByUUID(
		ctx context.Context,
		uuid string,
	) (Event, error)
	DeleteEvent(
		ctx context.Context,
		uuid string,
	) error
	IsMissingEventError(err error) bool
}

type Event struct {
	// UUID - уникальный идентификатор события.
	UUID string `db:"uuid"`
	// Заголовок - короткий текст.
	Summary string `db:"summary"`
	// Unix timestamp даты и времени начала события.
	StartedAt int32 `db:"started_at"`
	// Unix timestamp даты и времени завершения события.
	FinishedAt int32 `db:"finished_at"`
	// Описание события - длинный текст, опционально.
	Description string `db:"description"`
	// UUID пользователя, владельца события.
	UserUUID string `db:"user_uuid"`
	// Unix timestamp даты и времени уведомления о событии.
	NotificationAt int32 `db:"notification_at"`
}

type EventUpdateAttributes struct {
	// Заголовок - короткий текст.
	Summary string `db:"summary"`
	// Unix timestamp даты и времени начала события.
	StartedAt int32 `db:"started_at"`
	// Unix timestamp даты и времени завершения события.
	FinishedAt int32 `db:"finished_at"`
	// Описание события.
	Description string `db:"description"`
	// Get пользователя, владельца события.
	UserUUID string `db:"user_uuid"`
	// Unix timestamp даты и времени уведомления о событии.
	NotificationAt int32 `db:"notification_at"`
}
