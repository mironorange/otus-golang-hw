package main

import (
	"context"
	"fmt"
	sqlstorage "github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage/sql"
	"log"
)

// Пример приложения, которое использует возможности sql-хранилища.
func main() {
	driver := "postgres"
	dsn := "postgres://whoever:qwerty@localhost/test?sslmode=disable"
	storage := sqlstorage.New(driver, dsn)
	ctx := context.Background()
	if err := storage.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	// Создать в базе данных запись о событии
	e := sqlstorage.Event{
		UUID:           "a3970cf5-d262-431d-99dc-16bb95840032",
		Summary:        "Проверить создание события из main.go",
		StartedAt:      "2022-06-11 11:00:00",
		FinishedAt:     "2022-06-11 12:00:00",
		Description:    "После запуска скрипта это событие можно увидеть в базе данных",
		UserUUID:       "a6e592bc-8627-4e13-b4a6-d7072864602a",
		NotificationAt: "2022-06-11 10:00:00",
	}
	if _, err := storage.Create(e); err != nil {
		log.Fatal(err)
	}

	var uuid string
	
	// Обновить запись о событии в базе данных
	uuid = "a3970cf5-d262-431d-99dc-16bb95840032"
	attrs := sqlstorage.EventUpdateAttributes{
		Summary:        "Проверить изменения события из main.go",
		StartedAt:      "2022-06-21 11:00:00",
		FinishedAt:     "2022-06-21 12:00:00",
		Description:    "После запуска скрипта это событие должно измениться в базе данных",
		UserUUID:       "a6e592bc-8627-4e13-b4a6-d7072864602a",
		NotificationAt: "2022-06-21 10:00:00",
	}
	if _, err := storage.Update(uuid, attrs); err != nil {
		log.Fatal(err)
	}

	// Обновить запись о событии в базе данных
	uuid = "a3970cf5-d262-431d-99dc-16bb95840032"
	event, err := storage.Get(uuid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event)

	// Обновить запись о событии в базе данных
	events, err := storage.Select()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(events)
}
