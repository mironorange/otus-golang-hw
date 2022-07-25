package sqlstorage

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	// Необходимо импортировать пакет для того чтобы подключился драйвер pq.
	_ "github.com/lib/pq"
	"github.com/mironorange/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
)

var (
	schema = `
-- Схема для хранения событий в базе данных
CREATE SCHEMA IF NOT EXISTS "events";

-- Таблица для хранения событий в базе данных
CREATE TABLE IF NOT EXISTS "events"."events"
(
	-- Get - уникальный идентификатор события
    "uuid" varchar
		constraint events_pk
			primary key,
	-- Заголовок - короткий текст
	"summary" varchar not null,
	-- Unix timestamp даты и времени начала события.
	"started_at" int not null,
	-- Unix timestamp даты и времени завершения события.
	"finished_at" int not null,
	-- Описание события
	"description" varchar not null,
	-- Get пользователя, владельца события
	"user_uuid" varchar not null,
	-- Unix timestamp даты и времени уведомления о событии.
	"notification_at" int not null
);`

	sqlEventSelectByID = `SELECT * FROM "events"."events" WHERE "uuid" = $1 LIMIT 1`

	sqlEventDelete = `DELETE FROM "events"."events" WHERE "uuid" = $1 LIMIT 1`

	sqlGetEvents = `SELECT * FROM "events"."events"`

	sqlGetOldestEvents = `SELECT * FROM "events"."events" WHERE "finished_at" < $1`

	sqlGetEventsToBeNotified = `SELECT * FROM "events"."events" WHERE "notification_at" > $1 AND "notification_at" < $2`

	sqlEventInsert = `-- Запрос создающий запись в базе данных о событии 
INSERT INTO "events"."events"
(uuid, summary, started_at, finished_at, description, user_uuid, notification_at)
VALUES (:uuid, :summary, :started_at, :finished_at, :description, :user_uuid, :notification_at)`

	sqlEventUpdate = `-- Запрос обновляющий запись в базе данных о событии
UPDATE "events"."events"
SET summary = $2,
    started_at = $3,
    finished_at = $4,
    description = $5,
    user_uuid = $6,
    notification_at = $7
WHERE uuid = $1`
)

type Storage struct {
	driver    string
	dsn       string
	dbConnect *sqlx.DB
}

type EventStorage interface {
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
	storage.EventStorage
}

func New(driver string, dsn string) EventStorage {
	return &Storage{
		driver: driver,
		dsn:    dsn,
	}
}

func (s *Storage) Connect(ctx context.Context) error {
	db, err := sqlx.Connect(s.driver, s.dsn)
	if err == nil {
		db.MustExec(schema)
		s.dbConnect = db
	}

	return err
}

func (s *Storage) Close(ctx context.Context) error {
	if s.dbConnect != nil {
		return s.dbConnect.Close()
	}

	return nil
}

// Добавляет в базу данных новое событие.
func (s *Storage) CreateEvent(
	ctx context.Context,
	uuid string,
	summary string,
	startedAt int32,
	finishedAt int32,
	description string,
	userUUID string,
	notificationAt int32,
) error {
	event := storage.Event{
		UUID:           uuid,
		Summary:        summary,
		StartedAt:      startedAt,
		FinishedAt:     finishedAt,
		Description:    description,
		UserUUID:       userUUID,
		NotificationAt: notificationAt,
	}
	if _, err := s.dbConnect.NamedExec(sqlEventInsert, &event); err != nil {
		return err
	}
	return nil
}

// Обновляет существующее в хранилище событие.
func (s *Storage) UpdateEvent(
	ctx context.Context,
	uuid string,
	summary string,
	startedAt int32,
	finishedAt int32,
	description string,
	userUUID string,
	notificationAt int32,
) error {
	args := []interface{}{
		uuid,
		summary,
		startedAt,
		finishedAt,
		description,
		userUUID,
		notificationAt,
	}
	if _, err := s.dbConnect.Exec(sqlEventUpdate, args...); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetEvents(ctx context.Context, sinceNotificationAt int32) (events []storage.Event, err error) {
	err = s.dbConnect.Select(&events, sqlGetEvents)
	return events, err
}

func (s *Storage) GetEventsToBeNotified(ctx context.Context, from, to int32) (events []storage.Event, err error) {
	err = s.dbConnect.Select(&events, sqlGetEventsToBeNotified, from, to)
	return events, err
}

func (s *Storage) GetOldestEvents(ctx context.Context, endedAt int32) (events []storage.Event, err error) {
	err = s.dbConnect.Select(&events, sqlGetOldestEvents, endedAt)
	return events, err
}

func (s *Storage) GetEventByUUID(ctx context.Context, uuid string) (event storage.Event, err error) {
	err = s.dbConnect.Get(&event, sqlEventSelectByID, uuid)
	return event, err
}

func (s *Storage) IsMissingEventError(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

func (s *Storage) DeleteEvent(ctx context.Context, uuid string) error {
	_, err := s.dbConnect.Exec(sqlEventDelete, uuid)
	return err
}
