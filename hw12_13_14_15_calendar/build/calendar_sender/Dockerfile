# Собираем в гошке
FROM golang:1.16.2 as build

ENV BIN_FILE /opt/calendar/calendarsender
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

# Кэшируем слои с модулями
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY .. ${CODE_DIR}

# Собираем статический бинарник Go (без зависимостей на Си API),
# иначе он не будет работать в alpine образе.
ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/calendar_sender/*

# На выходе тонкий образ
FROM alpine:3.9

LABEL ORGANIZATION="OTUS Online Education"
LABEL SERVICE="calendar_sender"
LABEL MAINTAINERS="mironorange@gmail.com"

ENV BIN_FILE /opt/calendar/calendarsender
COPY --from=build ${BIN_FILE} ${BIN_FILE}

CMD ${BIN_FILE}
