package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	sqlstorage "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage/sql"
	// Необходимо импортировать пакет для того чтобы подключился драйвер pq.
	_ "github.com/lib/pq"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
)

var eventStorage sqlstorage.EventStorage

// Пример приложения, которое использует возможности sql-хранилища.
func main() {
	driver := "postgres"
	dsn := "postgres://whoever:qwerty@localhost/calendar?sslmode=disable"
	eventStorage = sqlstorage.New(driver, dsn)
	ctx := context.Background()
	if err := eventStorage.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	//iKnowHowCreateGetAndUpdate()
	//iKnowHowGetEventsToBeNotified()
	//iKnowHowGetOldestEvents()
	iKnowHowGetInstance()
}

func iKnowHowGetInstance() {
	driver := "postgres"
	dsn := "postgres://whoever:qwerty@localhost/calendar?sslmode=disable"
	db, _ := sqlx.Connect(driver, dsn)

	sqlGetEvent := `SELECT * FROM "events"."events" WHERE "uuid" = $1 LIMIT 1`

	event := storage.Event{}
	err := db.Get(&event, sqlGetEvent, "aa8e6639-fa2c-477d-a52c-0c7d64c19d75")

	fmt.Println(errors.Is(err, sql.ErrNoRows))
}

func iKnowHowCreateGetAndUpdate() {
	// Создать в базе данных запись о событии
	uid := uuid.NewString()
	ctx := context.Background()
	e := storage.Event{
		UUID:           uid,
		Summary:        "Проверить создание события из main.go",
		StartedAt:      1654070400,
		FinishedAt:     1654074000,
		Description:    "После запуска скрипта это событие можно увидеть в базе данных",
		UserUUID:       uuid.NewString(),
		NotificationAt: 1654070400,
	}
	err := eventStorage.CreateEvent(
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

	// Обновить запись о событии в базе данных
	attrs := storage.EventUpdateAttributes{
		Summary:        "Проверить изменения события из main.go",
		StartedAt:      1654243200,
		FinishedAt:     1654250400,
		Description:    "После запуска скрипта это событие должно измениться в базе данных",
		UserUUID:       uuid.NewString(),
		NotificationAt: 1654246800,
	}
	err = eventStorage.UpdateEvent(
		ctx,
		uid,
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
	event, err := eventStorage.GetEventByUUID(ctx, uid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event)

	// Обновить запись о событии в базе данных
	events, err := eventStorage.GetEvents(ctx, -1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(events))
}

func iKnowHowGetEventsToBeNotified() {
	ctx := context.Background()

	events, err := eventStorage.GetEventsToBeNotified(ctx, 1654060400, 1654080400)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(events))
}

func iKnowHowGetOldestEvents() {
	ctx := context.Background()

	events, err := eventStorage.GetOldestEvents(ctx, 1654060400)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(events))
}
