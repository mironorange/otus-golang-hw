package sqlstorage

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	schema = `
-- Схема для хранения событий в базе данных
CREATE SCHEMA IF NOT EXISTS "events";

-- Таблица для хранения событий в базе данных
CREATE TABLE IF NOT EXISTS "events"."events"
(
	-- UUID - уникальный идентификатор события
    "uuid" varchar
		constraint events_pk
			primary key,
	-- Заголовок - короткий текст
	"summary" varchar not null,
	-- Дата и время начала события
	"started_at" varchar not null,
	-- Дата и время начала события
	"finished_at" varchar not null,
	-- Описание события
	"description" varchar not null,
	-- UUID пользователя, владельца события
	"user_uuid" varchar not null,
	-- Дата и время уведомления о событии
	"notification_at" varchar not null
);`

	sqlEventSelectById = `SELECT * FROM "events"."events" WHERE "uuid" = $1 LIMIT 1`

	// Запрос создающий запись в базе данных о событии
	sqlEventInsert = `INSERT INTO "events"."events" (uuid, summary, started_at, finished_at, description, user_uuid, notification_at) VALUES (:uuid, :summary, :started_at, :finished_at, :description, :user_uuid, :notification_at)`

	// Запрос обновляющий запись в базе данных о событии
	sqlEventUpdate = `UPDATE "events"."events" SET summary = $2, started_at = $3, finished_at = $4, description = $5, user_uuid = $6, notification_at = $7 WHERE uuid = $1`
)

type Event struct {
	// UUID - уникальный идентификатор события.
	UUID string `db:"uuid"`
	// Заголовок - короткий текст.
	Summary string `db:"summary"`
	// Дата и время начала события.
	StartedAt string `db:"started_at"`
	// Дата и время начала события.
	FinishedAt string `db:"finished_at"`
	// Описание события.
	Description string `db:"description"`
	// UUID пользователя, владельца события.
	UserUUID string `db:"user_uuid"`
	// Дата и время уведомления о событии.
	NotificationAt string `db:"notification_at"`
}

type EventUpdateAttributes struct {
	// Заголовок - короткий текст.
	Summary string `db:"summary"`
	// Дата и время начала события.
	StartedAt string `db:"started_at"`
	// Дата и время начала события.
	FinishedAt string `db:"finished_at"`
	// Описание события.
	Description string `db:"description"`
	// UUID пользователя, владельца события.
	UserUUID string `db:"user_uuid"`
	// Дата и время уведомления о событии.
	NotificationAt string `db:"notification_at"`
}

type Storage struct {
	driver string
	dsn string
	dbConnect *sqlx.DB
}

type EventStorage interface {
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
	Create(attributes Event) (bool, error)
	Update(uuid string, attributes EventUpdateAttributes) (bool, error)
	Get() (map[string]Event, error)
	UUID(uuid string) (Event, error)
}

func New(driver string, dsn string) EventStorage {
	return &Storage{
		driver: driver,
		dsn: dsn,
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
func (s *Storage) Create(event Event) (bool, error) {
	tx := s.dbConnect.MustBegin()
	if _, err := tx.NamedExec(sqlEventInsert, &event); err != nil {
		_ = tx.Rollback()
		return false, err
	}
	if err := tx.Commit(); err != nil {
		return false, err
	}
	return true, nil
}

// Обновляет существующее в хранилище событие.
func (s *Storage) Update(uuid string, attributes EventUpdateAttributes) (bool, error) {
	tx := s.dbConnect.MustBegin()
	args := []interface{}{
		uuid,
		attributes.Summary,
		attributes.StartedAt,
		attributes.FinishedAt,
		attributes.Description,
		attributes.UserUUID,
		attributes.NotificationAt,
	}
	if _, err := tx.Exec(sqlEventUpdate, args...); err != nil {
		_ = tx.Rollback()
		return false, err
	}
	if err := tx.Commit(); err != nil {
		return false, err
	}
	return true, nil
}

// Возвращает список соответствующих условию событий из хранилища, проиндексированные по идентификатору.
func (s *Storage) Get() (map[string]Event, error) {
	events := map[string]Event{}
	items := []Event{}
	err := s.dbConnect.Select(&items, `SELECT * FROM "events"."events"`)
	if err != nil {
		return events, err
	}
	for _, event := range items {
		events[event.UUID] = event
	}
	return events, nil
}

// Возвращает событие из хранилища по идентификатору.
func (s *Storage) UUID(uuid string) (Event, error) {
	event := Event{}
	rows, err := s.dbConnect.Queryx(sqlEventSelectById, uuid)
	if err != nil {
		return Event{}, err
	}
	for rows.Next() {
		err = rows.StructScan(&event)
		if err != nil {
			return Event{}, err
		}
	}
	return event, nil
}
