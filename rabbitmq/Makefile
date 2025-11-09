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

NAME = ${RABBITMQ_CONTAINER}
PROJECT_NAME = ${RABBITMQ_CONTAINER}
SERVICE_NAME = ${RABBITMQ_CONTAINER}

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
# RABBITMQ
# ========================

# RABBITMQ = providers/rabbitmq

# .PHONY: rabbitmq-up rabbitmq-logs rabbitmq-stop rabbitmq-restart rabbitmq-rm rabbitmq-down rabbitmq-ps rabbitmq-exec

# rabbitmq-up:
# 	$(MAKE_CMD) -C $(RABBITMQ) up
# rabbitmq-logs:
# 	$(MAKE_CMD) -C $(RABBITMQ) logs
# rabbitmq-stop:
# 	$(MAKE_CMD) -C $(RABBITMQ) stop
# rabbitmq-restart:
# 	$(MAKE_CMD) -C $(RABBITMQ) restart
# rabbitmq-rm:
# 	$(MAKE_CMD) -C $(RABBITMQ) rm
# rabbitmq-down:
# 	$(MAKE_CMD) -C $(RABBITMQ) down
# rabbitmq-ps:
# 	$(MAKE_CMD) -C $(RABBITMQ) ps
# rabbitmq-exec:
# 	$(MAKE_CMD) -C $(RABBITMQ) exec