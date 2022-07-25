#!/bin/sh

set -xue

DB_NAME=$1

wait_for_db()
{
    for _ in $(seq 1 30); do
        psql -U "$POSTGRES_USER" -c "SELECT 1" && return
        sleep 1
    done

    exit 1
}

wait_for_db

# Создать базу данных, если она ранее не была создана
echo "SELECT 'CREATE DATABASE $DB_NAME' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '$DB_NAME')\gexec" \
  | psql -U "$POSTGRES_USER"

# Добавить к ней права от имени пользователя, с которым мы работаем с базой данных
echo "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $POSTGRES_USER" \
  | psql -U "$POSTGRES_USER"

# Создать схему если она ранее не была создана
echo "CREATE SCHEMA IF NOT EXISTS events" \
  | psql -U "$POSTGRES_USER" -d "$DB_NAME"

# Убедиться в том, что схема создана
echo "SELECT schema_name FROM information_schema.schemata WHERE schema_name = 'events'" \
  | psql -U "$POSTGRES_USER" -d "$DB_NAME"
