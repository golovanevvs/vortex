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

NAME = ${PROMTAIL_CONTAINER}
PROJECT_NAME = ${PROMTAIL_CONTAINER}
SERVICE_NAME = ${PROMTAIL_CONTAINER}

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
	@echo "Restarting Promtail..."

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
# PROMTAIL
# ========================

# PROMTAIL = providers/promtail

# .PHONY: promtail-up promtail-logs promtail-stop promtail-restart promtail-rm promtail-down promtail-ps promtail-exec

# promtail-up:
# 	$(MAKE_CMD) -C $(PROMTAIL) up
# promtail-logs:
# 	$(MAKE_CMD) -C $(PROMTAIL) logs
# promtail-stop:
# 	$(MAKE_CMD) -C $(PROMTAIL) stop
# promtail-restart:
# 	$(MAKE_CMD) -C $(PROMTAIL) restart
# promtail-rm:
# 	$(MAKE_CMD) -C $(PROMTAIL) rm
# promtail-down:
# 	$(MAKE_CMD) -C $(PROMTAIL) down
# promtail-ps:
# 	$(MAKE_CMD) -C $(PROMTAIL) ps
# promtail-exec:
# 	$(MAKE_CMD) -C $(PROMTAIL) exec