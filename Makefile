# === Env Load ===
ifneq (,$(wildcard ./.env))
include ./.env
export $(shell sed 's/=.*//' ./.env)
endif

# === Config ===
WORK_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

APP_NAME := calendar

DB_DSN := "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

MIGRATIONS_DIR := $(WORK_DIR)/migrations

GOOSE_REBUILD ?= 0
GOOSE_VERSION := v3.26.0
GOOSE_SRC := $(WORK_DIR)/tools/src/goose
GOOSE_BIN := $(WORK_DIR)/tools/bin/goose

# === Help ===
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  setup         -  full setup (docker, db, tools) (GOOSE_REBUILD=1 available)"
	@echo "  server        -  run server locally (RACE=1 available)"
	@echo "  scheduler     -  run scheduler locally (RACE=1 available)"
	@echo "  run-docker    -  start server and scheduler via docker compose"
	@echo "  setup-docker  -  build docker images"
	@echo "  setup-db      -  create database if not exists"
	@echo "  setup-goose   -  fetch and build custom goose binary (GOOSE_REBUILD=1 available)"
	@echo "  goose-up      -  apply migrations"
	@echo "  goose-down    -  rollback last migration"
	@echo "  goose-status  -  show migration status"
	@echo "  goose-create  -  create new migration (name=<migration_name>)"
	@echo "  env-check     -  validate required environment variables"

# === ENV Validations ===
REQUIRED_ENV := DB_USER DB_PASSWORD DB_HOST DB_PORT DB_NAME

.PHONY: env-check
env-check:
	@missing_vars=0; \
	for var in $(REQUIRED_ENV); do \
		val=$$(printenv $$var); \
		if [ -z "$$val" ]; then \
			echo "$(RED)Missing environment variable: $$var $(COLOR_END)"; \
			missing_vars=1; \
		fi; \
	done; \
	if [ $$missing_vars -ne 0 ]; then \
		exit 1; \
	else \
		echo "$(GREEN)==> ENV is valid$(COLOR_END)"; \
	fi


# === Setup ===
.PHONY: setup
setup: env-check setup-docker setup-db setup-goose
	@echo "$(GREEN)==> Setup complete$(COLOR_END)"

# === Docker ===
.PHONY: setup-docker
setup-docker:
	@echo "$(YELLOW)==> Building docker images$(COLOR_END)"
	docker compose build
	@echo "$(GREEN)==> Built docker images$(COLOR_END)"

# === Database ===
.PHONY: setup-db
setup-db:
	@echo "$(YELLOW)==> Ensuring database $(DB_NAME) exists$(COLOR_END)"
	docker compose exec -T db psql -U $(DB_USER) -tc "SELECT 1 FROM pg_database WHERE datname='$(DB_NAME)';" | grep -q 1 || \
	    docker compose exec -T db createdb -U $(DB_USER) $(DB_NAME)
	@echo "$(GREEN)==> Database ready$(COLOR_END)"

# === Goose ===
.PHONY: setup-goose
setup-goose:
	@echo "$(YELLOW)==> Setting up custom goose ($(GOOSE_VERSION))$(COLOR_END)"
ifeq ($(GOOSE_REBUILD),1)
	@rm -rf $(GOOSE_SRC) $(GOOSE_BIN)
endif
	@if [ ! -f "$(GOOSE_BIN)" ]; then \
		git clone -q --branch $(GOOSE_VERSION) https://github.com/pressly/goose.git $(GOOSE_SRC); \
		mkdir -p $(dir $(GOOSE_BIN)); \
		cd $(GOOSE_SRC) && go mod tidy; \
		cd $(GOOSE_SRC) && go build -tags='no_clickhouse no_libsql no_mssql no_mysql no_sqlite3 no_vertica no_ydb' \
			-o $(GOOSE_BIN) ./cmd/goose/; \
		rm -rf $(GOOSE_SRC); \
	fi
	@echo "$(GREEN)==> Built $(GOOSE_BIN)$(COLOR_END)"

# === Migrations ===
.PHONY: goose-up goose-down goose-status goose-create
goose-up:
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) postgres $(DB_DSN) up

goose-down:
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) postgres $(DB_DSN) down

goose-status:
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) postgres $(DB_DSN) status

goose-create:
ifndef name
	$(error name parameter required, usage: make goose-create name=create_users)
endif
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) create -s $(name) sql

# === Run ===
.PHONY: up-docker
uprun-docker:
	docker compose up -d app scheduler

.PHONY: server
server:
ifeq ($(RACE),1)
	go run -race ./cmd/server/server.go
else
	go run ./cmd/server/server.go
endif

.PHONY: scheduler
scheduler:
ifeq ($(RACE),1)
	go run -race ./cmd/scheduler/scheduler.go
else
	go run ./cmd/scheduler/scheduler.go
endif

# === Colors ===

GREEN := \033[0;32m
YELLOW := \033[1;33m
RED := \033[0;31m
COLOR_END := \033[0m
