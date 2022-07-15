package memorystorage

import (
	"context"
	"testing"

	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestCreateEvent(t *testing.T) {
	e := storage.Event{
		UUID:           "1753b11e-0841-4b6e-9d24-c21ea1b2d83d",
		Summary:        "Написать unit-тест и проверить создание события",
		StartedAt:      1654070400,
		FinishedAt:     1654074000,
		Description:    "Написанный unit-тест должен проходить без ошибок и может быть доработан",
		UserUUID:       "a6e592bc-8627-4e13-b4a6-d7072864602a",
		NotificationAt: 1654070400,
	}
	s := Storage{
		events: map[string]storage.Event{},
	}
	ctx := context.Background()

	// Создать объект события можно в случае, если событие ранее не было создано.
	err := s.CreateEvent(
		ctx,
		e.UUID,
		e.Summary,
		e.StartedAt,
		e.FinishedAt,
		e.Description,
		e.UserUUID,
		e.NotificationAt,
	)
	require.NoError(t, err)
	require.Less(t, 0, len(s.events))

	// Создать событие ранее было создано, то при создании возникнет ошибка.
	err = s.CreateEvent(
		ctx,
		e.UUID,
		e.Summary,
		e.StartedAt,
		e.FinishedAt,
		e.Description,
		e.UserUUID,
		e.NotificationAt,
	)
	require.Error(t, err)
	require.Equal(t, 1, len(s.events))
}

func TestUpdateEvent(t *testing.T) {
	uuid := "1753b11e-0841-4b6e-9d24-c21ea1b2d83d"
	event := storage.Event{
		UUID:           uuid,
		Summary:        "Написать unit-тест и проверить создание события",
		StartedAt:      1654070400,
		FinishedAt:     1654074000,
		Description:    "Написанный unit-тест должен проходить без ошибок и может быть доработан",
		UserUUID:       "a6e592bc-8627-4e13-b4a6-d7072864602a",
		NotificationAt: 1654070400,
	}
	eventUpdateAttributes := storage.EventUpdateAttributes{
		Summary:        "Изменить unit-тест и проверить создание события",
		StartedAt:      1654243200,
		FinishedAt:     1654250400,
		Description:    "В написанном тесте должно проверяться обновление события",
		UserUUID:       "a6e592bc-8627-4e13-b4a6-d7072864602a",
		NotificationAt: 1654246800,
	}
	s := Storage{
		events: map[string]storage.Event{},
	}
	ctx := context.Background()

	// Попытка обновить ранее не созданное событие приведет к ошибке.
	err := s.UpdateEvent(
		ctx,
		uuid,
		eventUpdateAttributes.Summary,
		eventUpdateAttributes.StartedAt,
		eventUpdateAttributes.FinishedAt,
		eventUpdateAttributes.Description,
		eventUpdateAttributes.UserUUID,
		eventUpdateAttributes.NotificationAt,
	)
	require.Error(t, err)

	// Создать объект события можно в случае, если событие ранее не было создано.
	err = s.CreateEvent(
		ctx,
		event.UUID,
		event.Summary,
		event.StartedAt,
		event.FinishedAt,
		event.Description,
		event.UserUUID,
		event.NotificationAt,
	)
	require.NoError(t, err)
	require.Less(t, 0, len(s.events))

	// Попытка обновить ранее существующее событие приведет пройдет без ошибок.
	err = s.UpdateEvent(
		ctx,
		uuid,
		eventUpdateAttributes.Summary,
		eventUpdateAttributes.StartedAt,
		eventUpdateAttributes.FinishedAt,
		eventUpdateAttributes.Description,
		eventUpdateAttributes.UserUUID,
		eventUpdateAttributes.NotificationAt,
	)
	require.NoError(t, err)

	event, err = s.GetEventByUUID(ctx, uuid)
	require.NoError(t, err)
	require.Equal(t, eventUpdateAttributes.Summary, event.Summary)
	require.Equal(t, eventUpdateAttributes.StartedAt, event.StartedAt)
	require.Equal(t, eventUpdateAttributes.FinishedAt, event.FinishedAt)
	require.Equal(t, eventUpdateAttributes.Description, event.Description)
	require.Equal(t, eventUpdateAttributes.UserUUID, event.UserUUID)
	require.Equal(t, eventUpdateAttributes.NotificationAt, event.NotificationAt)
}
