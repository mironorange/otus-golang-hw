BIN := "./bin/calendar"
SCHEDULER_BIN := "./bin/calendar_scheduler"
SENDER_BIN := "./bin/calendar_sender"
PB_EVENTS_PATH="./api/"
PB_GEN_PATH="./internal/pb"
DOCKER_IMG="calendar:develop"
DOCKER_SCHEDULER_IMG="calendarscheduler:develop"
DOCKER_SENDER_IMG="calendarsender:develop"
BROKER_QUEUE_URI="amqp://guest:guest@192.168.0.111:5672/"
EVENTSSERVICE_HOST="192.168.0.111"
EVENTSSERVICE_PORT=50051

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

ip-address:
	ifconfig en0 | grep "inet " | grep -v 127.0.0.1 | cut -d\  -f2

generate:
	[ -d "$(PB_GEN_PATH)" ] && rm -rf $(PB_GEN_PATH) && mkdir -p $(PB_GEN_PATH)
	protoc --proto_path=$(PB_EVENTS_PATH) --go_out=$(PB_GEN_PATH) --go-grpc_out=$(PB_GEN_PATH) $(PB_EVENTS_PATH)/EventService.proto

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

build-scheduler:
	go build -v -o $(SCHEDULER_BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar_scheduler

build-sender:
	go build -v -o $(SENDER_BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar_sender

run: build
	$(BIN) -config ./configs/config.json

run-scheduler: build-scheduler
	$(BIN) -config ./configs/config_scheduler.json

run-sender: build-sender
	$(BIN) -config ./configs/config_sender.json

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/calendar/Dockerfile .

build-img-scheduler:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_SCHEDULER_IMG) \
		-f build/calendar_scheduler/Dockerfile .

build-img-sender:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_SENDER_IMG) \
		-f build/calendar_sender/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

run-img-scheduler: build-img-scheduler
	docker run -e EVENTSSERVICE_PORT=$(EVENTSSERVICE_PORT) -e EVENTSSERVICE_HOST=$(EVENTSSERVICE_HOST) \
	    -e QUEUE_URI=$(BROKER_QUEUE_URI) $(DOCKER_SCHEDULER_IMG)

run-img-sender: build-img-sender
	docker run -e QUEUE_URI=$(BROKER_QUEUE_URI) \
	    $(DOCKER_SENDER_IMG)

version: build
	$(BIN) version

test:
	go test -race ./internal/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps
	golangci-lint run ./...

up:
    # calendar
    # scheduler
    # sender
    # rabbitmq
    # postgres
	docker-compose -f deployments/docker-compose.yaml up -d postgres rabbitmq
	bash deployments/scripts/dockerinitdb.sh
down:
	docker-compose -f deployments/docker-compose.yaml down

.PHONY: build run build-img run-img version test lint
