package sqlstorage

import (
	"context"

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
	"user_uuid" int not null,
	-- Unix timestamp даты и времени уведомления о событии.
	"notification_at" int not null
);`

	sqlEventSelectByID = `SELECT * FROM "events"."events" WHERE "uuid" = $1 LIMIT 1`

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
	tx := s.dbConnect.MustBegin()
	event := storage.Event{
		UUID:           uuid,
		Summary:        summary,
		StartedAt:      startedAt,
		FinishedAt:     finishedAt,
		Description:    description,
		UserUUID:       userUUID,
		NotificationAt: notificationAt,
	}
	if _, err := tx.NamedExec(sqlEventInsert, &event); err != nil {
		_ = tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
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
	tx := s.dbConnect.MustBegin()
	args := []interface{}{
		uuid,
		summary,
		startedAt,
		finishedAt,
		description,
		userUUID,
		notificationAt,
	}
	if _, err := tx.Exec(sqlEventUpdate, args...); err != nil {
		_ = tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// Возвращает список соответствующих условию событий из хранилища.
func (s *Storage) GetEvents(ctx context.Context) ([]storage.Event, error) {
	var events []storage.Event
	err := s.dbConnect.Select(&events, `SELECT * FROM "events"."events"`)
	return events, err
}

func (s *Storage) GetEventByUUID(ctx context.Context, uuid string) (storage.Event, error) {
	var event storage.Event
	rows, err := s.dbConnect.Queryx(sqlEventSelectByID, uuid)
	if err != nil {
		return storage.Event{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.StructScan(&event)
		if err != nil {
			return storage.Event{}, err
		}
	}
	return event, nil
}
