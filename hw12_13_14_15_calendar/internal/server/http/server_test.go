package internalhttp

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/logger"
	memorystorage "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage/memory"
	"github.com/stretchr/testify/require"
)

func TestCreateEvent(t *testing.T) {
	logging := logger.New("INFO")
	storage := memorystorage.New().(app.Storage)
	app := app.New(logging, storage)

	mux := http.NewServeMux()

	helloHandler := http.HandlerFunc(hello)
	eventsHandler := createEventsHandler(logging, app)
	mux.Handle("/events/", loggingMiddleware(eventsHandler, logging))
	mux.Handle("/", helloHandler)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	t.Run("Get main", func(t *testing.T) {
		request, _ := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			ts.URL,
			nil,
		)
		response, err := http.DefaultClient.Do(request)
		content, _ := io.ReadAll(response.Body)
		require.NoError(t, err)
		require.Equal(t, 200, response.StatusCode)
		require.Equal(t, []byte("Hello, World!"), content)
		response.Body.Close()
	})

	t.Run("Get unknown", func(t *testing.T) {
		request, _ := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			fmt.Sprintf("%s/any", ts.URL),
			nil,
		)
		response, err := http.DefaultClient.Do(request)
		require.NoError(t, err)
		require.Equal(t, 404, response.StatusCode)
		response.Body.Close()
	})

	t.Run("Create and update event", func(t *testing.T) {
		// Убедиться в том, что ранее такое событие не было создано.
		uuid := "1753b11e-0841-4b6e-9d24-c21ea1b2d83d"
		request, _ := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			fmt.Sprintf("%s/events/%s", ts.URL, uuid),
			nil,
		)
		response, err := http.DefaultClient.Do(request)
		require.NoError(t, err)
		require.Equal(t, 404, response.StatusCode)
		response.Body.Close()

		// Создать событие через REST API.
		exceptedEvent := Event{
			UUID:           uuid,
			Summary:        "Написать unit-тест на REST API и проверить создание события",
			StartedAt:      1654070400,
			FinishedAt:     1654074000,
			Description:    "Написанный unit-тест должен проходить без ошибок и может быть доработан",
			UserUUID:       "a6e592bc-8627-4e13-b4a6-d7072864602a",
			NotificationAt: 1654070400,
		}
		jsonBody, _ := exceptedEvent.MarshalJSON()
		request, _ = http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			fmt.Sprintf("%s/events/", ts.URL),
			bytes.NewBuffer(jsonBody),
		)
		response, err = http.DefaultClient.Do(request)
		require.NoError(t, err)
		require.Equal(t, 201, response.StatusCode)
		response.Body.Close()

		// Убедиться в том, что событие создано через REST API.
		resultEvent := Event{}
		request, _ = http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			fmt.Sprintf("%s/events/%s", ts.URL, uuid),
			nil,
		)
		response, err = http.DefaultClient.Do(request)
		content, _ := io.ReadAll(response.Body)
		_ = resultEvent.UnmarshalJSON(content)
		require.NoError(t, err)
		require.Equal(t, 200, response.StatusCode)
		require.Equal(t, exceptedEvent, resultEvent)
		response.Body.Close()

		// Обновить событие через REST API.
		exceptedUpdatedEvent := Event{
			UUID:           uuid,
			Summary:        "Обновленное через unit-тест на REST API событие",
			StartedAt:      1654080400,
			FinishedAt:     1654084000,
			Description:    "Обновленное описание",
			UserUUID:       "a6e592bc-8627-4e13-b4a6-7d072864602a",
			NotificationAt: 1654080400,
		}
		eventAttributes := EventUpdateAttributes{
			Summary:        exceptedUpdatedEvent.Summary,
			StartedAt:      exceptedUpdatedEvent.StartedAt,
			FinishedAt:     exceptedUpdatedEvent.FinishedAt,
			Description:    exceptedUpdatedEvent.Description,
			UserUUID:       exceptedUpdatedEvent.UserUUID,
			NotificationAt: exceptedUpdatedEvent.NotificationAt,
		}
		jsonBody, _ = eventAttributes.MarshalJSON()
		request, _ = http.NewRequestWithContext(
			context.Background(),
			http.MethodPut,
			fmt.Sprintf("%s/events/%s", ts.URL, uuid),
			bytes.NewBuffer(jsonBody),
		)
		response, err = http.DefaultClient.Do(request)
		require.NoError(t, err)
		require.Equal(t, 200, response.StatusCode)
		_ = response.Body.Close()

		// Убедиться в том, что событие в базе обновлено.
		resultEvent = Event{}
		request, _ = http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			fmt.Sprintf("%s/events/%s", ts.URL, uuid),
			nil,
		)
		response, err = http.DefaultClient.Do(request)
		content, _ = io.ReadAll(response.Body)
		_ = resultEvent.UnmarshalJSON(content)
		require.NoError(t, err)
		require.Equal(t, 200, response.StatusCode)
		require.Equal(t, exceptedUpdatedEvent, resultEvent)
		_ = response.Body.Close()
	})
}
