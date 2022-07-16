package app

import (
	"context"

	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
)

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
	) ([]storage.Event, error)
	GetOldestEvents(
		ctx context.Context,
		endedAt int32,
	) ([]storage.Event, error)
	GetEventByUUID(
		ctx context.Context,
		uuid string,
	) (storage.Event, error)
	DeleteEvent(
		ctx context.Context,
		uuid string,
	) error
}

type App struct {
	storage EventStorage
}

type Logger interface {
	Info(msg string)
}

func New(logger Logger, storage EventStorage) *App {
	return &App{
		storage: storage,
	}
}

func (a *App) CreateEvent(
	ctx context.Context,
	uuid string,
	summary string,
	startedAt int32,
	finishedAt int32,
	description string,
	userUUID string,
	notificationAt int32,
) error {
	return a.storage.CreateEvent(
		ctx,
		uuid,
		summary,
		startedAt,
		finishedAt,
		description,
		userUUID,
		notificationAt,
	)
}

func (a *App) UpdateEvent(
	ctx context.Context,
	uuid string,
	summary string,
	startedAt int32,
	finishedAt int32,
	description string,
	userUUID string,
	notificationAt int32,
) error {
	return a.storage.UpdateEvent(
		ctx,
		uuid,
		summary,
		startedAt,
		finishedAt,
		description,
		userUUID,
		notificationAt,
	)
}

func (a *App) GetEventByUUID(
	ctx context.Context,
	uuid string,
) (storage.Event, error) {
	return a.storage.GetEventByUUID(ctx, uuid)
}

func (a *App) GetEvents(
	ctx context.Context,
	sinceNotificationAt int32,
) ([]storage.Event, error) {
	return a.storage.GetEvents(ctx, sinceNotificationAt)
}

func (a *App) GetOldestEvents(
	ctx context.Context,
	endedAt int32,
) ([]storage.Event, error) {
	return a.storage.GetOldestEvents(ctx, endedAt)
}

func (a *App) DeleteEvent(
	ctx context.Context,
	uuid string,
) error {
	return a.storage.DeleteEvent(ctx, uuid)
}
