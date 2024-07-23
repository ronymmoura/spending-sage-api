include .env

MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
DB_URL := postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable

setup: create-db migrate-up

reset-db: drop-db setup

test:
	go test -cover -coverprofile=c.out ./...

test-coverage: |
	go test -cover -coverprofile=c.out ./...
	go tool cover -html=c.out -o coverage.html

server:
	go run main.go

# DB
create-db:
	docker exec -it postgres createdb --username=${DATABASE_USER} --owner=${DATABASE_USER} ${DATABASE_NAME}

drop-db:
	docker exec -it postgres dropdb ${DATABASE_NAME}

migration-new:
	migrate create -ext sql -dir=internal/db/migrations -seq $(name)

migrate-up:
	migrate -path=internal/db/migrations/ -database="$(DB_URL)" -verbose up

migrate-up1:
	migrate -path=internal/db/migrations/ -database="$(DB_URL)" -verbose up 1
	
migrate-down:
	migrate -path=internal/db/migrations/ -database="$(DB_URL)" -verbose down
	
migrate-down1:
	migrate -path=internal/db/migrations/ -database="$(DB_URL)" -verbose down 1


# SQLC
sqlc-init:
	docker run --rm -v $(MAKEFILE_DIR):/src -w /src sqlc/sqlc init

sqlc-compile:
	docker run --rm -v $(MAKEFILE_DIR):/src -w /src sqlc/sqlc compile -f /src/internal/db/sqlc.yaml

sqlc-generate:
	docker run --rm -v $(MAKEFILE_DIR):/src -w /src sqlc/sqlc generate -f /src/internal/db/sqlc.yaml

.PHONY:
	create-db drop-db migration-new migrate-up migrate-down migrate-up1 migrate-down1 