

export MY_WORKSPACE?=$(PWD)
export MY_USER?=root

OUTPUT_BASE?=./build

all: deps test build

build: build-server

clean: clean-output-dir

deps:
	go mod download

deps-fix:
	go mod tidy

check-fmt:
	test -z $(shell gofmt -l ./)

build-server: output-dir
	go build -o $(OUTPUT_BASE)/server acme/cmd/server

output-dir:
	mkdir -p $(OUTPUT_BASE)

#Comands for the use local
postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.1-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root acme

dropdb:
	docker exec -it postgres15 dropdb acme

migrateup:
	migrate --path migrations -database "postgresql://root:secret@localhost:5432/acme?sslmode=disable" -verbose up

migratedown:
	migrate --path migrations -database "postgresql://root:secret@localhost:5432/acme?sslmode=disable" -verbose down