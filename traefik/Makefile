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

NAME = ${TRAEFIK_CONTAINER}
PROJECT_NAME = ${TRAEFIK_CONTAINER}
SERVICE_NAME = ${TRAEFIK_CONTAINER}

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
# TRAEFIK
# ========================

# TRAEFIK = providers/traefik

# .PHONY: traefik-up traefik-logs traefik-stop traefik-restart traefik-rm traefik-down traefik-ps traefik-exec

# traefik-up:
# 	$(MAKE_CMD) -C $(TRAEFIK) up
# traefik-logs:
# 	$(MAKE_CMD) -C $(TRAEFIK) logs
# traefik-stop:
# 	$(MAKE_CMD) -C $(TRAEFIK) stop
# traefik-restart:
# 	$(MAKE_CMD) -C $(TRAEFIK) restart
# traefik-rm:
# 	$(MAKE_CMD) -C $(TRAEFIK) rm
# traefik-down:
# 	$(MAKE_CMD) -C $(TRAEFIK) down
# traefik-ps:
# 	$(MAKE_CMD) -C $(TRAEFIK) ps
# traefik-exec:
# 	$(MAKE_CMD) -C $(TRAEFIK) exec