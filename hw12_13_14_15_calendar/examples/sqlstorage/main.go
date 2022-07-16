package main

import (
	"context"
	"fmt"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
	sqlstorage "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage/sql"
	"log"
)

// Пример приложения, которое использует возможности sql-хранилища.
func main() {
	driver := "postgres"
	dsn := "postgres://whoever:qwerty@localhost/test?sslmode=disable"
	s := sqlstorage.New(driver, dsn)
	ctx := context.Background()
	if err := s.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	// Создать в базе данных запись о событии
	e := storage.Event{
		UUID:           "1753b11e-0841-4b6e-9d24-c21ea1b2d83d",
		Summary:        "Проверить создание события из main.go",
		StartedAt:      1654070400,
		FinishedAt:     1654074000,
		Description:    "После запуска скрипта это событие можно увидеть в базе данных",
		UserUUID:       "a6e592bc-8627-4e13-b4a6-d7072864602a",
		NotificationAt: 1654070400,
	}
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
	if err != nil {
		log.Fatal(err)
	}

	var uuid string
	
	// Обновить запись о событии в базе данных
	uuid = "a3970cf5-d262-431d-99dc-16bb95840032"
	attrs := storage.EventUpdateAttributes{
		Summary:        "Проверить изменения события из main.go",
		StartedAt:      1654243200,
		FinishedAt:     1654250400,
		Description:    "После запуска скрипта это событие должно измениться в базе данных",
		UserUUID:       "a6e592bc-8627-4e13-b4a6-d7072864602a",
		NotificationAt: 1654246800,
	}
	err = s.UpdateEvent(
		ctx,
		uuid,
		attrs.Summary,
		attrs.StartedAt,
		attrs.FinishedAt,
		attrs.Description,
		attrs.UserUUID,
		attrs.NotificationAt,
	)
	if  err != nil {
		log.Fatal(err)
	}

	// Обновить запись о событии в базе данных
	uuid = "a3970cf5-d262-431d-99dc-16bb95840032"
	event, err := s.GetEventByUUID(ctx, uuid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event)

	// Обновить запись о событии в базе данных
	events, err := s.GetEvents(ctx, -1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(events)
}
