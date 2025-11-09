MAKEFLAGS += --no-print-directory

PROJECT_DIR := $(abspath ../..)

ENV_FILE ?= .env ../traefik/.env
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

PROJECT_NAME = ${CLOUDPUB_CONTAINER}
NAME = ${CLOUDPUB_CONTAINER}
SERVICE_NAME = ${CLOUDPUB_CONTAINER}

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
# CLOUDPUB
# ========================

# CLOUDPUB = providers/cloudpub

# .PHONY: cloudpub-up cloudpub-logs cloudpub-stop cloudpub-restart cloudpub-rm cloudpub-down cloudpub-ps cloudpub-exec

# cloudpub-up:
# 	$(MAKE_CMD) -C $(CLOUDPUB) up
# cloudpub-logs:
# 	$(MAKE_CMD) -C $(CLOUDPUB) logs
# cloudpub-stop:
# 	$(MAKE_CMD) -C $(CLOUDPUB) stop
# cloudpub-restart:
# 	$(MAKE_CMD) -C $(CLOUDPUB) restart
# cloudpub-rm:
# 	$(MAKE_CMD) -C $(CLOUDPUB) rm
# cloudpub-down:
# 	$(MAKE_CMD) -C $(CLOUDPUB) down
# cloudpub-ps:
# 	$(MAKE_CMD) -C $(CLOUDPUB) ps
# cloudpub-exec:
# 	$(MAKE_CMD) -C $(CLOUDPUB) exec