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

NAME = ${REDIS_CONTAINER}
PROJECT_NAME = ${REDIS_CONTAINER}
SERVICE_NAME = ${REDIS_CONTAINER}

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
# REDIS
# ========================

# REDIS = providers/redis

# .PHONY: redis-up redis-logs redis-stop redis-restart redis-rm redis-down redis-ps redis-exec

# redis-up:
# 	$(MAKE_CMD) -C $(REDIS) up
# redis-logs:
# 	$(MAKE_CMD) -C $(REDIS) logs
# redis-stop:
# 	$(MAKE_CMD) -C $(REDIS) stop
# redis-restart:
# 	$(MAKE_CMD) -C $(REDIS) restart
# redis-rm:
# 	$(MAKE_CMD) -C $(REDIS) rm
# redis-down:
# 	$(MAKE_CMD) -C $(REDIS) down
# redis-ps:
# 	$(MAKE_CMD) -C $(REDIS) ps
# redis-exec:
# 	$(MAKE_CMD) -C $(REDIS) exec