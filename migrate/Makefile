MAKEFLAGS += --no-print-directory

PROJECT_DIR := $(abspath ../..)

ENV_FILE ?= $(PROJECT_DIR)/.env
ifneq (,$(wildcard $(ENV_FILE)))
include $(ENV_FILE)
export
endif

COMPOSE_FILE = docker-compose.yml
NETWORK_NAME = migrate_default

.PHONY: create up down

create:
	@echo "Creating new migration files..."
	docker compose -f $(COMPOSE_FILE) run --rm migrate create -ext sql -dir /migrations -seq tables
	@echo "Cleaning up leftover Docker network..."
	@docker network rm $(NETWORK_NAME) 2>/dev/null || true

up:
	@echo "Applying migrations..."
	docker compose -f $(COMPOSE_FILE) run --rm migrate up
	@docker network rm $(NETWORK_NAME) 2>/dev/null || true

down:
	@echo "Rollback all migrations..."
	docker compose -f $(COMPOSE_FILE) run --rm migrate down
	@docker network rm $(NETWORK_NAME) 2>/dev/null || true

#============================================================


# ========================
# MIGRATE
# ========================

# MIGRATE = providers/migrate

# .PHONY: migrate-create migrate-up migrate-down

# migrate-create:
# 	$(MAKE_CMD) -C $(MIGRATE) create
# migrate-up:
# 	$(MAKE_CMD) -C $(MIGRATE) up
# migrate-down:
# 	$(MAKE_CMD) -C $(MIGRATE) down