MAKEFLAGS += --no-print-directory

PROJECT_DIR := $(abspath ../..)

ENV_FILE ?= .env
ifneq (,$(wildcard $(ENV_FILE)))
include $(ENV_FILE)
export
endif

ifeq ($(OS),Windows_NT)
	MAKE_CMD = mingw32-make
else
	MAKE_CMD = make
endif

.PHONY: up logs stop restart rm down ps exec

NAME = ${LOKI_CONTAINER}
PROJECT_NAME = ${LOKI_CONTAINER}
SERVICE_NAME = ${LOKI_CONTAINER}

up:
	@echo "Starting ${NAME}..."
	docker compose -p $(PROJECT_NAME) up -d
	@echo "${NAME} is up and running."

logs:
	docker compose -p $(PROJECT_NAME) logs -f

stop:
	@echo "Stopping ${NAME}..."
	docker compose -p $(PROJECT_NAME) stop
	@echo "${NAME} stopped."

restart: stop up

rm:
	@echo "Removing ${NAME} (keeping data)..."
	docker compose -p $(PROJECT_NAME) rm -f
	@echo "${NAME} removed."

down:
	@echo "Full cleanup ${NAME}: containers, volumes, networks, data..."
	docker compose -p $(PROJECT_NAME) down -v --remove-orphans
	@echo "${NAME} cleaned."

ps:
	docker compose -p $(PROJECT_NAME) ps

exec:
	docker compose -p $(PROJECT_NAME) exec $(SERVICE_NAME) sh

# ========================
# LOKI
# ========================

# LOKI = providers/loki

# loki-up:
# 	$(MAKE_CMD) -C $(LOKI) up
# loki-logs:
# 	$(MAKE_CMD) -C $(LOKI) logs
# loki-stop:
# 	$(MAKE_CMD) -C $(LOKI) stop
# loki-restart:
# 	$(MAKE_CMD) -C $(LOKI) restart
# loki-rm:
# 	$(MAKE_CMD) -C $(LOKI) rm
# loki-down:
# 	$(MAKE_CMD) -C $(LOKI) down
# loki-ps:
# 	$(MAKE_CMD) -C $(LOKI) ps
# loki-exec:
# 	$(MAKE_CMD) -C $(LOKI) exec