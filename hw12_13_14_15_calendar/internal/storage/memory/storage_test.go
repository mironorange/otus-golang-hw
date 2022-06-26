package memorystorage

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestCreateEvent(t *testing.T) {
	e := Event{
		UUID:           "1753b11e-0841-4b6e-9d24-c21ea1b2d83d",
		Summary:        "Написать unit-тест и проверить создание события",
		StartedAt:      "2022-06-01 11:00:00",
		FinishedAt:     "2022-06-01 12:00:00",
		Description:    "Написанный unit-тест должен проходить без ошибок и может быть доработан",
		UserUUID:       "a6e592bc-8627-4e13-b4a6-d7072864602a",
		NotificationAt: "2022-06-01 10:00:00",
	}
	s := Storage{
		events: map[string]Event{},
	}

	// Создать объект события можно в случае, если событие ранее не было создано
	ok, err := s.Create(e)
	require.NoError(t, err)
	require.True(t, ok)
	require.Less(t, 0, len(s.events))

	// Создать событие ранее было создано, то при создании возникнет ошибка
	ok, err = s.Create(e)
	require.Error(t, err)
	require.False(t, ok)
	require.Equal(t, 1, len(s.events))
}

func TestUpdateEvent(t *testing.T) {
	uuid := "1753b11e-0841-4b6e-9d24-c21ea1b2d83d"
	event := Event{
		UUID:           uuid,
		Summary:        "Написать unit-тест и проверить создание события",
		StartedAt:      "2022-06-01 11:00:00",
		FinishedAt:     "2022-06-01 12:00:00",
		Description:    "Написанный unit-тест должен проходить без ошибок и может быть доработан",
		UserUUID:       "a6e592bc-8627-4e13-b4a6-d7072864602a",
		NotificationAt: "2022-06-01 10:00:00",
	}
	eventUpdateAttributes := EventUpdateAttributes{
		Summary:        "Изменить unit-тест и проверить создание события",
		StartedAt:      "2022-06-03 11:00:00",
		FinishedAt:     "2022-06-03 12:00:00",
		Description:    "В написанном тесте должно проверяться обновление события",
		UserUUID:       "a6e592bc-8627-4e13-b4a6-d7072864602a",
		NotificationAt: "2022-06-03 10:00:00",
	}
	s := Storage{
		events: map[string]Event{},
	}

	// Попытка обновить ранее не созданное событие приведет к ошибке
	ok, err := s.Update(uuid, eventUpdateAttributes)
	require.Error(t, err)
	require.False(t, ok)

	// Создать объект события можно в случае, если событие ранее не было создано
	ok, err = s.Create(event)
	require.NoError(t, err)
	require.True(t, ok)
	require.Less(t, 0, len(s.events))

	// Попытка обновить ранее существующее событие приведет пройдет без ошибок
	ok, err = s.Update(uuid, eventUpdateAttributes)
	require.NoError(t, err)
	require.True(t, ok)

	event, err = s.UUID(uuid)
	require.NoError(t, err)
	require.Equal(t, eventUpdateAttributes.Summary, event.Summary)
	require.Equal(t, eventUpdateAttributes.StartedAt, event.StartedAt)
	require.Equal(t, eventUpdateAttributes.FinishedAt, event.FinishedAt)
	require.Equal(t, eventUpdateAttributes.Description, event.Description)
	require.Equal(t, eventUpdateAttributes.UserUUID, event.UserUUID)
	require.Equal(t, eventUpdateAttributes.NotificationAt, event.NotificationAt)
}
